package main

import (
	"fmt"
	"log"
)

func main() {
	input := `{
		"id": "d5475d58-9a02-49ae-ba2e-93d6a6c2e39e",
		"name": "iPhone 6s",
		"price": 649.99,
		"isAvailable": true,
		"colors" : ["White", "Grey"],
		"discount" : null
	}`

	tokens, err := Tokenize(input)

	if err != nil {
		log.Fatal("Tokenization Error: ", err)
	}
	printTokens(tokens)

	node, err := Parser(tokens)

	if err != nil {
		log.Fatal("Parser Error: ", err)
	}

	fmt.Printf("\n\nAST Node: %v\n", node)
}
