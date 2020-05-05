package main

import (
	"fmt"
	"os"
)

func main() {
	line, _ := os.LookupEnv("LINE")
	fmt.Println(line)
}
