package ring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	capacity := int64(-1)
	buf, err := New[int](capacity)
	assert.Nil(t, buf, "ring/New():\nwant  nil\ngot  %+v", buf)
	assert.EqualError(t, err, ErrBufferCapacity.Error(), "ring/New() error:\nwant  %+v\ngot  %+v", ErrBufferCapacity.Error(), err)

	capacity = 4
	buf, err = New[int](capacity)
	expected := &RingBuffer[int]{
		buffer:   make([]int, capacity),
		head:     0,
		write:    0,
		size:     0,
		capacity: capacity,
	}
	assert.Equal(t, expected, buf, "ring/New():\nwant  %+v\ngot  %+v", expected, buf)
	assert.Nil(t, err, "ring/New() error:\nwant  nil\ngot  %+v", err)
}

func TestPut(t *testing.T) {
	tests := []struct {
		capacity     int64
		putItems     []int
		expectedSize int64
	}{
		{
			capacity:     1,
			putItems:     []int{1},
			expectedSize: 1,
		},
		{
			capacity:     5,
			putItems:     []int{1, 2, 3},
			expectedSize: 3,
		},
		{
			capacity:     2,
			putItems:     []int{1, 2, 3, 4, 5},
			expectedSize: 2,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("Test case #%d", idx), func(t *testing.T) {
			buf, err := New[int](tt.capacity)
			assert.Nil(t, err, "ring/New() error:\nwant  nil\ngot  %+v", err)
			assert.Equal(t, tt.capacity, buf.Capacity(), "capacity:\nwant  %+v\ngot  %+v", tt.capacity, buf.Capacity())

			for _, v := range tt.putItems {
				buf.Put(v)
			}

			size := buf.Size()
			assert.Equal(t, tt.expectedSize, size, "size:\nwant  %+v\ngot  %+v", tt.expectedSize, size)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		capacity              int64
		putItems              []int
		expectedSizeBeforeGet int64
		expectedGetValues     []int
		expectedSizeAfterGet  int64
		errorString           string
	}{
		{
			capacity:              1,
			putItems:              []int{1},
			expectedSizeBeforeGet: 1,
			expectedGetValues:     []int{1},
			expectedSizeAfterGet:  0,
		},
		{
			capacity:              5,
			putItems:              []int{1, 2, 3},
			expectedSizeBeforeGet: 3,
			expectedGetValues:     []int{1, 2},
			expectedSizeAfterGet:  1,
		},
		{
			capacity:              2,
			putItems:              []int{1, 2, 3, 4, 5},
			expectedSizeBeforeGet: 2,
			expectedGetValues:     []int{4, 5},
			expectedSizeAfterGet:  0,
		},
		{
			capacity:              1,
			putItems:              []int{},
			expectedSizeBeforeGet: 0,
			expectedGetValues:     []int{1},
			errorString:           ErrBufferEmpty.Error(),
			expectedSizeAfterGet:  0,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("Test case #%d", idx), func(t *testing.T) {
			buf, err := New[int](tt.capacity)
			assert.Nil(t, err, "ring/New() error:\nwant  nil\ngot  %+v", err)
			assert.Equal(t, tt.capacity, buf.Capacity(), "capacity:\nwant  %+v\ngot  %+v", tt.capacity, buf.Capacity())

			for _, v := range tt.putItems {
				buf.Put(v)
			}

			size := buf.Size()
			assert.Equal(t, tt.expectedSizeBeforeGet, size, "before get size:\nwant  %+v\ngot  %+v", tt.expectedSizeBeforeGet, size)

			for _, v := range tt.expectedGetValues {
				value, err := buf.Get()
				if err != nil {
					assert.EqualError(t, err, tt.errorString, "ring/Get() error:\nwant  %+v\ngot  %+v", tt.errorString, err)
				} else {
					assert.Equal(t, v, value, "ring/Get():\nwant  %+v\ngot  %+v", v, value)
				}
			}

			size = buf.Size()
			assert.Equal(t, tt.expectedSizeAfterGet, size, "after get size:\nwant  %+v\ngot  %+v", tt.expectedSizeAfterGet, size)
		})
	}
}

func TestSize(t *testing.T) {
	buf, err := New[int](4)
	assert.Nil(t, err, "ring/New() error:\nwant  nil\ngot  %+v", err)

	// when buf is empty
	output := buf.Size()
	expected := int64(0)
	assert.Equal(t, expected, output, "size:\nwant  %+v\ngot  %+v", expected, output)

	// after putting two values
	buf.Put(12)
	buf.Put(10)
	output = buf.Size()
	expected = int64(2)
	assert.Equal(t, expected, output, "size:\nwant  %+v\ngot  %+v", expected, output)

	// after getting one value
	buf.Get()
	output = buf.Size()
	expected = int64(1)
	assert.Equal(t, expected, output, "size:\nwant  %+v\ngot  %+v", expected, output)
}

func TestCapacity(t *testing.T) {
	capacity := int64(4)
	buf, err := New[int](capacity)
	assert.Nil(t, err, "ring/New() error:\nwant  nil\ngot  %+v", err)
	output := buf.Capacity()
	assert.Equal(t, capacity, output, "capacity:\nwant  %+v\ngot  %+v", capacity, output)
}
