package main

import (
	"fmt"
	"log"
	"os"
	"transaction_processor/internal"
	"transaction_processor/utils"
)

func main() {
	// Set the configurable variables default
	bufferSize := 100

	var config utils.Config
	if err := config.InitConfig(); err != nil {
		log.Fatal("Config is not loaded application wont start")
	}

	dataFile := "resources/data.json"

	// Open the data file
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Printf("Failed to open data file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the ring buffer
	ringBuffer := internal.NewRingBuffer(bufferSize)

	// Create the transaction processor
	processor := internal.NewTransactionProcessor(ringBuffer)

	// Process the transactions
	err = processor.Process(file)
	if err != nil {
		fmt.Printf("Failed to process transactions: %v\n", err)
	}

	// Just to test locally
	transaction := ringBuffer.ReadAll()
	fmt.Println("Transaction ", transaction)
}
