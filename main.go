package main

import (
	"fmt"
	"log"
)

func main() {
	input := `{ "user": { "id": 1024, "name": "Sushant \"The Dev\"", "bio": "Backend Engineer", "active": true, "roles": ["admin", "moderator", "editor"], "preferences": { "notifications": { "email": true, "sms": false }, "theme": null } }, "projects": [ { "id": "p001", "name": "RateShield", "techStack": ["Go", "Redis", "REST"], "contributors": 3 }, { "id": "p002", "name": "LogPump", "techStack": ["Rust", "Elasticsearch", "TCP"], "contributors": 2 } ], "lastLogin": "2025-05-25T14:30:00Z", "metadata": null }`

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
