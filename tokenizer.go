package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
	Example JSON

	{
		"id": "d5475d58-9a02-49ae-ba2e-93d6a6c2e39e",
		"name": "iPhone 6s",
		"price": 649.99,
		"isAvailable": true,
		"colors" : ["White", "Grey"],
		"discount" : null
	}

*/

const (
	TKN_BRACE_OPEN    = "BRACE_OPEN"
	TKN_BRACE_CLOSE   = "BRACE_CLOSE"
	TKN_STRING        = "STRING"
	TKN_NUMBER        = "NUMBER"
	TKN_COLON         = "COLON"
	TKN_COMMA         = "COMMA"
	TKN_TRUE          = "TRUE"
	TKN_FALSE         = "FALSE"
	TKN_NULL          = "NULL"
	TKN_BRACKET_OPEN  = "BRACKET_OPEN"
	TKN_BRACKET_CLOSE = "BRACKET_CLOSE"
)

type Token struct {
	Type  string
	Value string
}

func Tokenize(jsonString string) ([]Token, error) {
	current := 0
	stringLength := len(jsonString)

	tokens := []Token{}

	for current < stringLength {
		char := jsonString[current]

		if unicode.IsSpace(rune(char)) {
			current++
			continue
		}

		switch char {
		case '{':
			tokens = append(tokens, Token{Type: TKN_BRACE_OPEN, Value: "{"})
			current++
		case '}':
			tokens = append(tokens, Token{Type: TKN_BRACE_CLOSE, Value: "}"})
			current++
		case '[':
			tokens = append(tokens, Token{Type: TKN_BRACKET_OPEN, Value: "["})
			current++
		case ']':
			tokens = append(tokens, Token{Type: TKN_BRACKET_CLOSE, Value: "]"})
			current++
		case ':':
			tokens = append(tokens, Token{Type: TKN_COLON, Value: ":"})
			current++
		case ',':
			tokens = append(tokens, Token{Type: TKN_COMMA, Value: ","})
			current++
		case '"':
			current++
			temp := current
			str := ""

			for jsonString[temp] != '"' {
				str += string(jsonString[temp])
				temp++
			}

			tokens = append(tokens, Token{Type: TKN_STRING, Value: str})

			current = temp
			current++
			continue

		default:
			rest := jsonString[current:]

			if strings.HasPrefix(rest, "true") {
				tokens = append(tokens, Token{Type: TKN_TRUE, Value: "true"})
				current += 4
			} else if strings.HasPrefix(rest, "false") {
				tokens = append(tokens, Token{Type: TKN_FALSE, Value: "false"})
				current += 5
			} else if strings.HasPrefix(rest, "null") {
				tokens = append(tokens, Token{Type: TKN_NULL, Value: "null"})
				current += 4
			} else if unicode.IsNumber(rune(char)) || char == '-' {
				start := current
				current++

				for current < stringLength && (unicode.IsNumber(rune(jsonString[current])) || jsonString[current] == '.') {
					current++
				}

				number := jsonString[start:current]

				if _, err := strconv.ParseFloat(number, 64); err != nil {
					return nil, fmt.Errorf("invalid number: %s", number)
				}

				tokens = append(tokens, Token{Type: TKN_NUMBER, Value: number})
			} else {
				return nil, fmt.Errorf("unexpected character: %c", char)
			}
		}
	}
	return tokens, nil
}

func printTokens(tokens []Token) {
	fmt.Printf("%-14s | %s\n", "Type", "Value")
	fmt.Println(strings.Repeat("-", 50))

	for _, token := range tokens {
		fmt.Printf("%-14s | %s\n", token.Type, token.Value)
	}
}
