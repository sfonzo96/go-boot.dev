package main

import "fmt"

/*
	Pointers are basically, variables that store memory address's, typically from another variable
	var p *int is the used syntax to define a pointer (blank with nil value). There's a more common way to do it:

	myString := "hello"
	myStringPtr := &myString // The & operator returns the memory address of the variable at its right, i.e. generates a pointer

	If you were to change the underlying value that the memory address referenced by a pointer stores, you should use the asterisc operator, it gives access to the value :

	*myStringPts = "new hello"
*/
type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// Don't touch below this line

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")
}
