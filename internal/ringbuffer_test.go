package internal

import (
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	rb := NewRingBuffer(5)

	if rb.size != 0 {
		t.Errorf("Expected size to be 0, got %d", rb.size)
	}

	if rb.capacity != 5 {
		t.Errorf("Expected capacity to be 5, got %d", rb.capacity)
	}

	if len(rb.buffer) != 5 {
		t.Errorf("Expected data length to be 5, got %d", len(rb.buffer))
	}

	// Ensure head and tail are initialized properly
	if rb.head != 0 || rb.tail != 0 {
		t.Errorf("Expected head and tail to be 0, got %d and %d", rb.head, rb.tail)
	}
}

func TestWriteAndReadAll(t *testing.T) {
	rb := NewRingBuffer(3)

	data := []Transaction{
		{
			ID:        "1",
			Value:     "300",
			Timestamp: "08-25-2023",
		}, {
			ID:        "2",
			Value:     "200",
			Timestamp: "08-25-2023",
		},
	}
	for _, d := range data {
		rb.Write(d)
	}

	result := rb.ReadAll()
	if len(result) != len(data) {
		t.Errorf("Expected result length to be %d, got %d", len(data), len(result))
	}

	for i, r := range result {
		if r != data[i] {
			t.Errorf("Expected %v at index %d, got %v", data[i], i, r)
		}
	}
}

func TestReadByID(t *testing.T) {
	rb := NewRingBuffer(3)
	rb.Write(Transaction{
		ID:        "100",
		Value:     "200",
		Timestamp: "08-25-2023",
	})
	rb.Write(Transaction{
		ID:        "200",
		Value:     "200",
		Timestamp: "08-25-2023",
	})

	_, err := rb.ReadByID(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = rb.ReadByID(3)
	if err == nil {
		t.Errorf("Expected an error for invalid ID")
	}
}

func TestDeleteByID(t *testing.T) {
	rb := NewRingBuffer(3)
	rb.Write(Transaction{
		ID:        "100",
		Value:     "200",
		Timestamp: "08-25-2023",
	})

	err := rb.DeleteByID(100)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result := rb.ReadAll()
	expected := []Transaction{
		{
			ID:        "100",
			Value:     "200",
			Timestamp: "08-25-2023",
		},
	}
	if len(result) != len(expected) {
		t.Errorf("Expected result length to be %d, got %d", len(expected), len(result))
	}

	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, r)
		}
	}

	err = rb.DeleteByID(3)
	if err == nil {
		t.Errorf("Expected an error for invalid ID")
	}
}
