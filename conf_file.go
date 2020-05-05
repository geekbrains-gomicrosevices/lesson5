package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	cfgPath := flag.String("config", "default.json", "Path to config file")
	flag.Parse()

	content, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := struct {
		Line string
	}{}
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Line)
}
