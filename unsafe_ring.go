package ring

type UnsafeRingBuffer[T any] struct {
	*RingBuffer[T]
}

// UnsafeNew creates a new instance of thread unsafe ring buffer.
func UnsafeNew[T any](capacity int64) (*UnsafeRingBuffer[T], error) {
	if capacity <= 0 {
		return nil, ErrBufferCapacity
	}

	return &UnsafeRingBuffer[T]{RingBuffer: new[T](capacity)}, nil
}

// Put adds a new element to ring buffer.
// If the ring buffer is already full,
// the oldest element will be overwritten.
func (r *UnsafeRingBuffer[T]) Put(value T) {
	r.insertValue(value)
}

// Get and remove the oldest element.
// If the ring buffer is empty, Get() will return error.
func (r *UnsafeRingBuffer[T]) Get() (T, error) {
	if r.isEmpty() {
		return getZero[T](), ErrBufferEmpty
	}
	return r.removeValue(), nil
}

// Size returns current size of ring buffer.
func (r UnsafeRingBuffer[T]) Size() int64 {
	return r.size
}

// Capacity returns capacity of ring buffer.
func (r UnsafeRingBuffer[T]) Capacity() int64 {
	return r.capacity
}
