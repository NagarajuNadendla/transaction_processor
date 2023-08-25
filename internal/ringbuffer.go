package internal

import (
	"errors"
	"sync"
)

// RingBuffer is a capacity-bounded circular queue
type RingBuffer struct {
	buffer   []Transaction
	capacity int
	size     int
	writeIdx int
	readIdx  int
	head     int
	tail     int
	mu       sync.Mutex
}

// NewRingBuffer creates a new RingBuffer with the given capacity
func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		buffer:   make([]Transaction, capacity),
		size:     0,
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// Write adds a transaction to the RingBuffer
func (rb *RingBuffer) Write(transaction Transaction) error {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.size == rb.capacity {
		return errors.New("RingBuffer is full")
	}

	rb.buffer[rb.tail] = transaction
	rb.tail = (rb.tail + 1) % rb.capacity
	rb.size++

	return nil

}

// ReadAll retrieves and removes a transaction from the RingBuffer
func (rb *RingBuffer) ReadAll() []interface{} {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	result := make([]interface{}, rb.size)
	idx := rb.head
	for i := 0; i < rb.size; i++ {
		result[i] = rb.buffer[idx]
		idx = (idx + 1) % rb.capacity
	}
	return result
}

func (rb *RingBuffer) ReadByID(id int) (Transaction, error) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if id < 0 || id >= rb.size {
		return Transaction{}, errors.New("Invalid ID")
	}

	idx := (rb.head + id) % rb.capacity
	return rb.buffer[idx], nil
}

// DeleteByID Delete transaction by id
func (rb *RingBuffer) DeleteByID(id int) error {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if id < 0 || id >= rb.size {
		return errors.New("Invalid ID")
	}

	idx := (rb.head + id) % rb.capacity
	rb.size--

	// Shift elements to close the gap
	for i := idx; i != rb.tail; i = (i + 1) % rb.capacity {
		next := (i + 1) % rb.capacity
		rb.buffer[i] = rb.buffer[next]
	}

	rb.tail = (rb.tail - 1 + rb.capacity) % rb.capacity
	return nil
}
