package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to ask a question and get the user's response
func askQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Function to calculate BMI
func calculateBMI(weight float64, height float64) float64 {
	// Convert height from cm to meters
	heightMeters := height / 100
	return weight / (heightMeters * heightMeters)
}

// Function to suggest health advice based on inputs
func suggestHealthAdvice(age int, bmi float64, conditions string, symptoms string, activityLevel string) {
	fmt.Println("\nHealth Advice Based on Your Information:")

	// Advice based on age
	if age >= 65 {
		fmt.Println("• As you are 65 or older, regular checkups and preventive care are important.")
	} else if age >= 45 {
		fmt.Println("• At your age, it's important to monitor your cardiovascular health and manage any risk factors.")
	} else {
		fmt.Println("• Maintaining an active and healthy lifestyle at your age will benefit you in the long term.")
	}

	// Advice based on BMI
	if bmi < 18.5 {
		fmt.Printf("• Your BMI is %.2f, which is considered underweight. A balanced diet and strength training may help.\n", bmi)
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Printf("• Your BMI is %.2f, which is considered normal. Keep maintaining your healthy lifestyle!\n", bmi)
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Printf("• Your BMI is %.2f, which is considered overweight. Consider a balanced diet and regular exercise.\n", bmi)
	} else {
		fmt.Printf("• Your BMI is %.2f, which is considered obese. Consult with a healthcare provider for a personalized plan.\n", bmi)
	}

	// Advice based on pre-existing conditions
	if strings.Contains(strings.ToLower(conditions), "diabetes") {
		fmt.Println("• Since you mentioned having diabetes, it's important to manage your blood sugar levels carefully.")
	}
	if strings.Contains(strings.ToLower(conditions), "hypertension") {
		fmt.Println("• Managing your blood pressure is essential due to your hypertension. Consider a heart-healthy diet.")
	}

	// Advice based on symptoms
	if strings.Contains(strings.ToLower(symptoms), "fever") || strings.Contains(strings.ToLower(symptoms), "cough") {
		fmt.Println("• Since you're experiencing fever or cough, consider getting checked for infections or respiratory issues.")
	} else if strings.Contains(strings.ToLower(symptoms), "pain") {
		fmt.Println("• Pain symptoms should not be ignored. It's advisable to consult a healthcare professional.")
	}

	// Advice based on activity level
	switch strings.ToLower(activityLevel) {
	case "sedentary":
		fmt.Println("• A sedentary lifestyle can increase health risks. Consider incorporating more physical activity into your routine.")
	case "active":
		fmt.Println("• Great! Being active is beneficial for your overall health. Keep it up!")
	case "very active":
		fmt.Println("• A very active lifestyle is excellent for maintaining health, but ensure you're getting enough recovery.")
	default:
		fmt.Println("• Staying active is important, no matter your current activity level.")
	}
}

func main() {
	fmt.Println("Welcome to the Health AI Agent")

	// Asking basic health questions
	name := askQuestion("What is your name?")
	ageStr := askQuestion("How old are you?")
	weightStr := askQuestion("What is your weight (in kg)?")
	heightStr := askQuestion("What is your height (in cm)?")
	conditions := askQuestion("Do you have any pre-existing medical conditions? (e.g., diabetes, hypertension)")

	// Additional health-related questions
	symptoms := askQuestion("Are you currently experiencing any symptoms? (e.g., cough, fever, pain)")
	activityLevel := askQuestion("How would you describe your physical activity level? (e.g., sedentary, active, very active)")

	// Convert age, weight, and height to numeric values
	age, _ := strconv.Atoi(ageStr)
	weight, _ := strconv.ParseFloat(weightStr, 64)
	height, _ := strconv.ParseFloat(heightStr, 64)

	// Calculate BMI
	bmi := calculateBMI(weight, height)

	// Summary
	fmt.Println("\nThank you for your responses. Here is a summary of your answers:")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %s\n", age)
	fmt.Printf("Weight: %s kg\n", weight)
	fmt.Printf("Height: %s cm\n", height)
	fmt.Printf("Pre-existing conditions: %s\n", conditions)
	fmt.Printf("Current symptoms: %s\n", symptoms)
	fmt.Printf("Physical activity level: %s\n", activityLevel)

	// Suggest health advice based on inputs
	suggestHealthAdvice(age, bmi, conditions, symptoms, activityLevel)

	fmt.Println("\nBased on your answers, our AI agent will suggest some health advice soon!")
}
