package ring

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	capacity := int64(-1)
	buf, err := New[int](capacity)
	assert.Nil(t, buf, "New():\nwant  nil\ngot  %+v", buf)
	assert.EqualError(t, err, ErrBufferCapacity.Error(), "New() error:\nwant  %+v\ngot  %+v", ErrBufferCapacity.Error(), err)

	capacity = 4
	buf, err = New[int](capacity)
	expected := &SafeRingBuffer[int]{
		RingBuffer: new[int](capacity),
	}
	assert.Equal(t, expected, buf, "New():\nwant  %+v\ngot  %+v", expected, buf)
	assert.Nil(t, err, "New() error:\nwant  nil\ngot  %+v", err)
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
			assert.Nil(t, err, "New() error:\nwant  nil\ngot  %+v", err)
			assert.Equal(t, tt.capacity, buf.Capacity(), "capacity:\nwant  %+v\ngot  %+v", tt.capacity, buf.Capacity())

			var wg sync.WaitGroup
			for _, v := range tt.putItems {
				wg.Add(1)
				go func(value int) {
					defer wg.Done()
					buf.Put(value)
				}(v)
			}
			wg.Wait()

			size := buf.Size()
			assert.Equal(t, tt.expectedSize, size, "size:\nwant  %+v\ngot  %+v", tt.expectedSize, size)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		capacity          int64
		putItems          int
		getCount          int
		expectedGetValues int
		errorString       string
	}{
		{
			capacity:          1,
			putItems:          1,
			getCount:          2,
			expectedGetValues: 1,
		},
		{
			capacity:          5,
			putItems:          3,
			getCount:          2,
			expectedGetValues: 2,
		},
		{
			capacity:          10,
			putItems:          3,
			getCount:          6,
			expectedGetValues: 3,
		},
		{
			capacity:          3,
			putItems:          1,
			getCount:          4,
			expectedGetValues: 1,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("Test case #%d", idx), func(t *testing.T) {
			buf, err := New[int](tt.capacity)
			assert.Nil(t, err, "New() error:\nwant  nil\ngot  %+v", err)
			assert.Equal(t, tt.capacity, buf.Capacity(), "capacity:\nwant  %+v\ngot  %+v", tt.capacity, buf.Capacity())

			for i := 0; i < tt.putItems; i++ {
				buf.Put(i)
			}

			var wg sync.WaitGroup
			dataChan := make(chan int, tt.putItems)
			for i := 0; i < tt.getCount; i++ {
				wg.Add(1)
				go func(value int) {
					defer wg.Done()
					value, err := buf.Get()
					if err == nil {
						dataChan <- value
					}
				}(i)
			}

			wg.Wait()
			close(dataChan)

			var count int
			for range dataChan {
				count += 1
			}

			assert.Equal(t, tt.expectedGetValues, count, "Get() values count:\nwant  %+v\ngot  %+v", tt.expectedGetValues, count)
		})
	}
}

func TestSize(t *testing.T) {
	buf, err := New[int](4)
	assert.Nil(t, err, "New() error:\nwant  nil\ngot  %+v", err)

	output := buf.Size()
	expected := int64(0)
	assert.Equal(t, expected, output, "size:\nwant  %+v\ngot  %+v", expected, output)

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			buf.Put(value)
		}(i)
	}

	go func() {
		defer wg.Done()
		buf.Get()
	}()

	wg.Wait()

	output = buf.Size()
	expected = int64(4)
	assert.Equal(t, expected, output, "size:\nwant  %+v\ngot  %+v", expected, output)
}

func TestCapacity(t *testing.T) {
	capacity := int64(4)
	buf, err := New[int](capacity)
	assert.Nil(t, err, "New() error:\nwant  nil\ngot  %+v", err)
	output := buf.Capacity()
	assert.Equal(t, capacity, output, "capacity:\nwant  %+v\ngot  %+v", capacity, output)
}
