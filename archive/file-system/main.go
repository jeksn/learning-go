package main

import (
	"fmt"
	// "io" 	// This was added from ChatGPT and serves no purpose, keeping it here for reference!
	"os"
)

func main() {
	// Create a new file
	fileName := "example.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// Write data to the file
	data := []byte("This is the example content!")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	// // This was added from ChatGPT and serves no purpose, keeping it here for reference!

	// Read data from the file
	// readData := make([]byte, len(data))
	// _, err = file.Read(readData)
	// if err != nil && err != io.EOF {
	// 	fmt.Println("Failed to read file:", err)
	// 	return
	// }

	fmt.Println("File contents:", string(data))
}
