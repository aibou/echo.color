package main

import "fmt"
import "strings"
import "os"

func main() {
	fmt.Printf("\033[32m%s\033[0m\n", strings.Join(os.Args[1:], " "))
}
