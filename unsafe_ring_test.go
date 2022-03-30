package ring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsafeNew(t *testing.T) {
	capacity := int64(-1)
	buf, err := UnsafeNew[int](capacity)
	assert.Nil(t, buf, "UnsafeNew():\nwant  nil\ngot  %+v", buf)
	assert.EqualError(t, err, ErrBufferCapacity.Error(), "UnsafeNew() error:\nwant  %+v\ngot  %+v", ErrBufferCapacity.Error(), err)

	capacity = 4
	buf, err = UnsafeNew[int](capacity)
	expected := &UnsafeRingBuffer[int]{
		RingBuffer: new[int](capacity),
	}
	assert.Equal(t, expected, buf, "UnsafeNew():\nwant  %+v\ngot  %+v", expected, buf)
	assert.Nil(t, err, "UnsafeNew() error:\nwant  nil\ngot  %+v", err)
}

func TestPutUnsafe(t *testing.T) {
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
			buf, err := UnsafeNew[int](tt.capacity)
			assert.Nil(t, err, "UnsafeNew() error:\nwant  nil\ngot  %+v", err)
			assert.Equal(t, tt.capacity, buf.Capacity(), "capacity:\nwant  %+v\ngot  %+v", tt.capacity, buf.Capacity())

			for _, v := range tt.putItems {
				buf.Put(v)
			}

			size := buf.Size()
			assert.Equal(t, tt.expectedSize, size, "size:\nwant  %+v\ngot  %+v", tt.expectedSize, size)
		})
	}
}

func TestGetUnsafe(t *testing.T) {
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
			buf, err := UnsafeNew[int](tt.capacity)
			assert.Nil(t, err, "UnsafeNew() error:\nwant  nil\ngot  %+v", err)
			assert.Equal(t, tt.capacity, buf.Capacity(), "capacity:\nwant  %+v\ngot  %+v", tt.capacity, buf.Capacity())

			for _, v := range tt.putItems {
				buf.Put(v)
			}

			size := buf.Size()
			assert.Equal(t, tt.expectedSizeBeforeGet, size, "before get size:\nwant  %+v\ngot  %+v", tt.expectedSizeBeforeGet, size)

			for _, v := range tt.expectedGetValues {
				value, err := buf.Get()
				if err != nil {
					assert.EqualError(t, err, tt.errorString, "Get() error:\nwant  %+v\ngot  %+v", tt.errorString, err)
				} else {
					assert.Equal(t, v, value, "Get():\nwant  %+v\ngot  %+v", v, value)
				}
			}

			size = buf.Size()
			assert.Equal(t, tt.expectedSizeAfterGet, size, "after get size:\nwant  %+v\ngot  %+v", tt.expectedSizeAfterGet, size)
		})
	}
}

func TestSizeUnsafe(t *testing.T) {
	buf, err := UnsafeNew[int](4)
	assert.Nil(t, err, "UnsafeNew() error:\nwant  nil\ngot  %+v", err)

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

func TestCapacityUnsafe(t *testing.T) {
	capacity := int64(4)
	buf, err := UnsafeNew[int](capacity)
	assert.Nil(t, err, "UnsafeNew() error:\nwant  nil\ngot  %+v", err)
	output := buf.Capacity()
	assert.Equal(t, capacity, output, "capacity:\nwant  %+v\ngot  %+v", capacity, output)
}
