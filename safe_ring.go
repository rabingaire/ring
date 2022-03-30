package ring

import "sync"

type SafeRingBuffer[T any] struct {
	*RingBuffer[T]
	mu sync.Mutex
}

// New creates a new instance of thread safe ring buffer.
func New[T any](capacity int64) (*SafeRingBuffer[T], error) {
	if capacity <= 0 {
		return nil, ErrBufferCapacity
	}

	return &SafeRingBuffer[T]{RingBuffer: new[T](capacity)}, nil
}

// Put adds a new element to ring buffer.
// If the ring buffer is already full,
// the oldest element will be overwritten.
func (r *SafeRingBuffer[T]) Put(value T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.insertValue(value)
}

// Get and remove the oldest element.
// If the ring buffer is empty, Get() will return error.
func (r *SafeRingBuffer[T]) Get() (T, error) {
	if r.isEmpty() {
		return getZero[T](), ErrBufferEmpty
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.removeValue(), nil
}

// Size returns current size of ring buffer.
func (r *SafeRingBuffer[T]) Size() int64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.size
}

// Capacity returns capacity of ring buffer.
func (r *SafeRingBuffer[T]) Capacity() int64 {
	return r.capacity
}
