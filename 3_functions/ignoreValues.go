package main

import "fmt"

func main() {
	// getNames returns two values, but I'm not using the second one
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}
