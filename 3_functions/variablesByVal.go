package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	// Have to re-asign the new value, otherwise as variables are passed by value the value stored in the memory adress wont change
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
	fmt.Scanln()
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd

	return sendsSoFar
}
