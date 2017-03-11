package main

import 	(
	"fmt"
	"os"
	"strings"
)
func main() {
	args := os.Args
	searchCriteria := strings.Join(args[1:], " ")
	fmt.Println(searchCriteria)
}


// 1) We learned that a go program's executable part must be part of "package main". Saying "go run"
// will compile it and run the compiled file

// 2) We learned that to make the CLI available as a command (ie "$> goshy"), we need to build the source and
// output it to a file the same name as the desired command. Then to run the executable without giving the full
// path (ie ./goshy), do export PATH=$PATH:. from within the directory where goshy's main file is.

// 3) We learned that we can access positional argmuments from the array returned by os.Args

// 4) We learned that we have to use Google Custom search for searching for images. We need to create a developer account
// and generate an API key