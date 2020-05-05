package main

import (
	"flag"
	"fmt"
)

func main() {
	strPointer := flag.String("line", "default_line", "Line to print")
	flag.Parse()
	fmt.Println(*strPointer)
}
