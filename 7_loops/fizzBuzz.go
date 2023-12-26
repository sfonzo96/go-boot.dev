package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i % 5 == 0 && i % 3 == 0 {
			fmt.Println("fizzbuzz");
			continue;
		}

		if i % 3 == 0 {
			fmt.Println("fizz")
			continue;
		}

		if i % 5 == 0 {
			fmt.Println("buzz")
			continue;
		}

		fmt.Println(i)
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
	fmt.Scanln()
}
