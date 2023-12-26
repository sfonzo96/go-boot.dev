package main

import "fmt"

// If two or more params are the same type I dont need to specify every type
func concat(s1, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
	fmt.Scanln()
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
	
}
