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

type ringBuffer struct {
	buffer   []interface{}
	head     int64
	write    int64
	size     int64
	capacity int64
}

// New creates a new instance of ring buffer.
func New(capacity int64) (*ringBuffer, error) {
	if capacity <= 0 {
		return nil, ErrBufferCapacity
	}

	return &ringBuffer{
		buffer:   make([]interface{}, capacity),
		head:     0,
		write:    0,
		size:     0,
		capacity: capacity,
	}, nil
}

// Put adds a new element to ring buffer.
// If the ring buffer is already full,
// the oldest element will be overwritten.
func (r *ringBuffer) Put(value interface{}) {
	r.insertValue(value)
}

// Get and remove the oldest element.
// If the ring buffer is empty, Get() will return error.
func (r *ringBuffer) Get() (interface{}, error) {
	if r.isEmpty() {
		return nil, ErrBufferEmpty
	}
	return r.removeValue(), nil
}

// Size returns current size of ring buffer.
func (r ringBuffer) Size() int64 {
	return r.size
}

// Capacity returns capacity of ring buffer.
func (r ringBuffer) Capacity() int64 {
	return r.capacity
}

// updateHead advances head, will snap around if goes over bound.
func (r *ringBuffer) updateHead() {
	r.head = ((r.head + 1) % r.capacity)
}

// updateWrite advances write, will snap around if goes over bound.
func (r *ringBuffer) updateWrite() {
	r.write = ((r.write + 1) % r.capacity)
}

// removeValue removes value, advances head, decreases size and return removed value.
func (r *ringBuffer) removeValue() interface{} {
	value := r.buffer[r.head]
	r.buffer[r.head] = nil
	r.size -= 1
	r.updateHead()
	return value
}

// insertValue will set value, checks if ring buffer is full,
// if its full it advances head, if not it increases size.
// finally it will advance write.
func (r *ringBuffer) insertValue(value interface{}) {
	r.buffer[r.write] = value
	if r.isFull() {
		r.updateHead()
	} else {
		r.size += 1
	}
	r.updateWrite()
}

// isFull checks if ring buffer is full.
func (r ringBuffer) isFull() bool {
	return r.size >= r.capacity
}

// isEmpty checks if ring buffer is empty.
func (r ringBuffer) isEmpty() bool {
	return r.size == 0
}
