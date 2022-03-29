package ring

import (
	"errors"
)

// Error constants for ring package
var (
	ErrBufferCapacity = errors.New("buffer capacity must be greater than zero")
	ErrBufferEmpty    = errors.New("buffer is empty")
)

type Ring interface {
	Put(value interface{})
	Get() interface{}
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

// New creates a new instance of ring buffer.
func New[T any](capacity int64) (*RingBuffer[T], error) {
	if capacity <= 0 {
		return nil, ErrBufferCapacity
	}

	return &RingBuffer[T]{
		buffer:   make([]T, capacity),
		head:     0,
		write:    0,
		size:     0,
		capacity: capacity,
	}, nil
}

// Put adds a new element to ring buffer.
// If the ring buffer is already full,
// the oldest element will be overwritten.
func (r *RingBuffer[T]) Put(value T) {
	r.insertValue(value)
}

// Get and remove the oldest element.
// If the ring buffer is empty, Get() will return error.
func (r *RingBuffer[T]) Get() (T, error) {
	if r.isEmpty() {
		return getZero[T](), ErrBufferEmpty
	}
	return r.removeValue(), nil
}

// Size returns current size of ring buffer.
func (r RingBuffer[T]) Size() int64 {
	return r.size
}

// Capacity returns capacity of ring buffer.
func (r RingBuffer[T]) Capacity() int64 {
	return r.capacity
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
