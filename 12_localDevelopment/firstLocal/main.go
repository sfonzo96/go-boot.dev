/*
	In order to start a go module/project this should be written in the terminal:
	"go mod init githubRepoLink"

	Then, "go run filename" will run a file as a little script.
	If you'd want to create an executable file that contains the whole module it's to be created with:
	"go build" from the module directory.
*/

package main

import (
	"fmt"

	"github.com/sfonzo96/mystring"
)

func main() {
	fmt.Println(mystring.Reverse("hello world"))
}
