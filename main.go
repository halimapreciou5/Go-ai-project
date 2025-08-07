package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if user provided any arguments
	if len(os.Args) < 2 {
		fmt.Println("AI CLI Tool")
		fmt.Println("Usage: ai-cli [prompt]")
		fmt.Println("Example: ai-cli \"Tell me a joke\"")
		return
	}

	// Get the prompt from command line arguments
	prompt := os.Args[1]
	
	fmt.Printf("You entered: %s\n", prompt)
	fmt.Println("AI response will go here...")
	
	// TODO: Add AI API integration
}