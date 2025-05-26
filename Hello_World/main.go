package main

import (
	"flag"
	"fmt"
	"hello_world/greeter"
	"log"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "Language code (supported: en, fr, es, la)")
	flag.Parse()

	greeting, err := greeter.Greet(greeter.Language(lang))
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	fmt.Println(greeting)
}
