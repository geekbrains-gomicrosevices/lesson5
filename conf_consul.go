package main

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"log"
)

func main() {

	cons, err := consul.NewClient(consul.DefaultConfig())

	if err != nil {
		log.Fatal(err)
	}

	line, _, err := cons.KV().Get("line", nil)
	if err != nil {
		log.Fatal(err)
	} else if line == nil {
		log.Print("Value isn't set")
	} else {
		fmt.Println(string(line.Value))
	}
}
