package main

import (
	"fmt"

	"github.com/google/uuid"
)

// GenerateUUID returns a new random UUID string.
func GenerateUUID() string {
	return uuid.New().String()
}

func main() {
	fmt.Printf("Generated UUID: %s\n", GenerateUUID())
}
