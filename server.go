package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Simulate an I/O operation (reading a file)
	file, err := os.Open("test.txt")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	io.Copy(w, file) // Send file content as response

	duration := time.Since(start)
	fmt.Printf("Handled request in %v\n", duration)
}

func Main() {
	http.HandleFunc("/", handler)
	fmt.Println("Go server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
