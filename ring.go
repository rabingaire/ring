package ring

import (
	"errors"
)

// Error constants for ring package
var (
	ErrBufferCapacity = errors.New("buffer capacity must be greater than zero")
	ErrBufferEmpty    = errors.New("buffer is empty")
)

type Ring[T any] interface {
	Put(value T)
	Get() T
	Size() int64
	Capacity() int64
}

type RingBuffer[T any] struct {
	buffer   []T
	head     int64
	write    int64
	size     int64
	capacity int64
}

// new returns a new instance of RingBuffer
func new[T any](capacity int64) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer:   make([]T, capacity),
		capacity: capacity,
	}
}

// updateHead advances head, will snap around if goes over bound.
func (r *RingBuffer[T]) updateHead() {
	r.head = ((r.head + 1) % r.capacity)
}

// updateWrite advances write, will snap around if goes over bound.
func (r *RingBuffer[T]) updateWrite() {
	r.write = ((r.write + 1) % r.capacity)
}

// removeValue removes value, advances head, decreases size and return removed value.
func (r *RingBuffer[T]) removeValue() T {
	value := r.buffer[r.head]
	r.buffer[r.head] = getZero[T]()
	r.size -= 1
	r.updateHead()
	return value
}

// insertValue will set value, checks if ring buffer is full,
// if its full it advances head, if not it increases size.
// finally it will advance write.
func (r *RingBuffer[T]) insertValue(value T) {
	r.buffer[r.write] = value
	if r.isFull() {
		r.updateHead()
	} else {
		r.size += 1
	}
	r.updateWrite()
}

// isFull checks if ring buffer is full.
func (r RingBuffer[T]) isFull() bool {
	return r.size >= r.capacity
}

// isEmpty checks if ring buffer is empty.
func (r RingBuffer[T]) isEmpty() bool {
	return r.size == 0
}

// getZero returns zero value for given type generic type T
func getZero[T any]() T {
	var result T
	return result
}
