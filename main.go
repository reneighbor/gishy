package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	searchCriteria := strings.Join(os.Args[1:], " ")
	fmt.Println("You searched for " + searchCriteria)
}


// 1) We learned that a go program's executable part must be part of "package 
//main". Saying "go run [directory]" will  compile package main and run the 
//compiled file 

// 2) We learned that to make the CLI available as a command (ie "$> goshy"), 
// need to build build the source and output it to a file the same name as
// the desired command. So "go build -o goshy"

// to make the command available without having to prepend with the path, add
// the current dir to the path; "export PATH=$PATH:$PWD"

// 3) We learned that we can access positional arguments using os.Args, which
// returns an array

// 4) We learned that you have to use Google Custom Search to search for images. 
// We need to create a developer account and generate an API key.