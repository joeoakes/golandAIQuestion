package golandAIQuestion

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

package main

import (
"bufio"
"fmt"
"os"
"strings"
)

// Function to ask a question and get the user's response
func askQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

func main() {
	fmt.Println("Welcome to the Health AI Agent")

	// Asking basic health questions
	name := askQuestion("What is your name?")
	age := askQuestion("How old are you?")
	weight := askQuestion("What is your weight (in kg)?")
	height := askQuestion("What is your height (in cm)?")
	conditions := askQuestion("Do you have any pre-existing medical conditions? (e.g., diabetes, hypertension)")

	// Additional health-related questions
	symptoms := askQuestion("Are you currently experiencing any symptoms? (e.g., cough, fever, pain)")
	activityLevel := askQuestion("How would you describe your physical activity level? (e.g., sedentary, active, very active)")

	// Summary
	fmt.Println("\nThank you for your responses. Here is a summary of your answers:")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %s\n", age)
	fmt.Printf("Weight: %s kg\n", weight)
	fmt.Printf("Height: %s cm\n", height)
	fmt.Printf("Pre-existing conditions: %s\n", conditions)
	fmt.Printf("Current symptoms: %s\n", symptoms)
	fmt.Printf("Physical activity level: %s\n", activityLevel)

	fmt.Println("\nBased on your answers, our AI agent will suggest some health advice soon!")
}
