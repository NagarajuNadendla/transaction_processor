package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// Transaction represents a transaction record
type Transaction struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

// TransactionProcessor processes the transactions from the data feed
type TransactionProcessor struct {
	ringBuffer *RingBuffer
}

// NewTransactionProcessor creates a new TransactionProcessor with the given ring buffer
func NewTransactionProcessor(ringBuffer *RingBuffer) *TransactionProcessor {
	return &TransactionProcessor{
		ringBuffer: ringBuffer,
	}
}

// Process starts reading and processing the transactions
func (tp *TransactionProcessor) Process(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		transaction := Transaction{}
		err := json.Unmarshal(line, &transaction)
		if err != nil {
			return err
		}
		tp.ringBuffer.Write(transaction)
		tp.processTransaction(transaction)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// processTransaction processes a single transaction
func (tp *TransactionProcessor) processTransaction(transaction Transaction) {
	// Placeholder code: print the transaction record to stdout
	fmt.Printf("Processed transaction: %+v\n", transaction)
}
