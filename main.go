package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Check if user provided any arguments
	if len(os.Args) < 2 {
		fmt.Println("🤖 AI CLI Tool (Mock Version)")
		fmt.Println("Usage: go run main.go \"[your prompt]\"")
		fmt.Println("Example: go run main.go \"Tell me a joke\"")
		return
	}

	// Get the prompt from command line arguments
	prompt := os.Args[1]
	
	fmt.Printf("🔍 Processing: %s\n", prompt)
	fmt.Println("⏳ Generating AI response...")
	
	// Simulate processing time
	time.Sleep(2 * time.Second)

	// Get AI response
	response := getMockAIResponse(prompt)
	fmt.Printf("\n🤖 AI Response:\n%s\n", response)
}

func getMockAIResponse(prompt string) string {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	promptLower := strings.ToLower(prompt)
	
	// Different responses based on prompt content
	if strings.Contains(promptLower, "joke") {
		jokes := []string{
			"Why don't scientists trust atoms? Because they make up everything!",
			"Why did the programmer quit his job? He didn't get arrays!",
			"Why do programmers prefer dark mode? Because light attracts bugs!",
		}
		return jokes[rand.Intn(len(jokes))]
	}
	
	if strings.Contains(promptLower, "hello") || strings.Contains(promptLower, "hi") {
		greetings := []string{
			"Hello! How can I help you today?",
			"Hi there! What would you like to know?",
			"Greetings! I'm here to assist you.",
		}
		return greetings[rand.Intn(len(greetings))]
	}
	
	if strings.Contains(promptLower, "weather") {
		return "I'm a mock AI, so I can't check real weather, but I hope it's nice where you are!"
	}
	
	if strings.Contains(promptLower, "time") {
		return fmt.Sprintf("The current time is: %s", time.Now().Format("3:04 PM"))
	}
	
	if strings.Contains(promptLower, "programming") || strings.Contains(promptLower, "code") {
		tips := []string{
			"Programming tip: Always write comments for your future self!",
			"Remember: The best code is code that doesn't need to exist.",
			"Debug tip: If it works, don't touch it... until you have to.",
		}
		return tips[rand.Intn(len(tips))]
	}
	
	// Default responses for other prompts
	defaults := []string{
		"That's an interesting question! As a mock AI, I can only give you pre-programmed responses.",
		"I understand you're asking about: " + prompt + ". This is a demonstration AI response.",
		"Thank you for your input! In a real AI system, this would be processed by machine learning models.",
		"Interesting prompt! This mock AI is showing how the CLI tool works without needing API keys.",
	}
	
	return defaults[rand.Intn(len(defaults))]
}