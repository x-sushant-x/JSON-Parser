package main

import "fmt"

func main() {
	input := `{ "user": { "id": 1024, "name": "Sushant \"The Dev\"", "bio": "Backend Engineer", "active": true, "roles": ["admin", "moderator", "editor"], "preferences": { "notifications": { "email": true, "sms": false }, "theme": null } }, "projects": [ { "id": "p001", "name": "RateShield", "techStack": ["Go", "Redis", "REST"], "contributors": 3 }, { "id": "p002", "name": "LogPump", "techStack": ["Rust", "Elasticsearch", "TCP"], "contributors": 2 } ], "lastLogin": "2025-05-25T14:30:00Z", "metadata": null }`

	tokens, err := Tokenize(input)

	if err != nil {
		fmt.Println("Tokenization Error: ", err)
	} else {
		printTokens(tokens)
	}
}
