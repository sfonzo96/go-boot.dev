/*
Complete the waitForDbs function. It should block until it receives numDBs tokens on the dbChan channel. Each time it reads a token the getDatabasesChannel goroutine will print a message to the console for you.

Tokens are values that are used in order to check if something happened or not, for ex. a channel can return an empty struct when a thread ended executing some code. The channel coll will block the code until a value is returned.
*/

package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<- dbChan
	} 
}

// don't touch below this line

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(3)
	test(4)
	test(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}
