package main

import "fmt"

func main() {
	input := `{ "id": "d5475d58-9a02-49ae-ba2e-93d6a6c2e39e", "name": "iPhone 6s", "price": 649.99, "isAvailable": true, "colors" : ["White", "Grey"], "discount" : null }`

	tokens, err := Tokenize(input)

	if err != nil {
		fmt.Println("Tokenization Error: ", err)
	} else {
		printTokens(tokens)
	}
}
