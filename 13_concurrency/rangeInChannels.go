/*
Complete the concurrentFib function. It should:

- Create a new channel of ints
- Call fibonacci in a goroutine, passing it the channel and the number of Fibonacci numbers to generate, n
- Use a range loop to read from the channel and print out the numbers one by one, each on a new line
*/

package main

import (
	"fmt"
	"time"
)

func concurrentFib(n int) {
	intCh := make(chan int)

	go fibonacci(n, intCh)

	for number := range intCh {
		fmt.Println(number)
	}
}

// don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
