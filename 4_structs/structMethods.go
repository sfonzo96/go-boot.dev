package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// Asignes a method to a struct type (extensions methods like feature)
func (a authenticationInfo) getBasicAuth() (username, password string) {
	return a.username, a.password
}
// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
	fmt.Scanln()
}
