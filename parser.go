package main

import (
	"errors"
	"fmt"
	"strconv"
)

type ASTNode interface {
	Type() string
}

type ObjectNode struct {
	Value map[string]ASTNode
}

func (o ObjectNode) Type() string {
	return "Object"
}

type ArrayNode struct {
	Value []ASTNode
}

func (a ArrayNode) Type() string {
	return "Array"
}

type StringNode struct {
	Value string
}

func (s StringNode) Type() string {
	return "String"
}

type NumberNode struct {
	Value float64
}

func (n NumberNode) Type() string {
	return "Number"
}

type BooleanNode struct {
	Value bool
}

func (b BooleanNode) Type() string {
	return "Boolean"
}

type NullNode struct{}

func (n NullNode) Type() string {
	return "Null"
}

func Parser(tokens []Token) (ASTNode, error) {
	if len(tokens) == 0 {
		return nil, errors.New("nothing to parse")
	}

	current := 0

	return parseValue(&current, tokens)
}

func parseValue(current *int, tokens []Token) (ASTNode, error) {
	if *current >= len(tokens) {
		return nil, fmt.Errorf("unexpected end of input")
	}

	token := tokens[*current]

	switch token.Type {
	case TKN_STRING:
		*current++
		return StringNode{Value: token.Value}, nil

	case TKN_NUMBER:
		num, _ := strconv.ParseFloat(token.Value, 64)
		*current++
		return NumberNode{Value: num}, nil

	case TKN_TRUE:
		*current++
		return BooleanNode{Value: true}, nil

	case TKN_FALSE:
		*current++
		return BooleanNode{Value: false}, nil

	case TKN_NULL:
		*current++
		return NullNode{}, nil

	case TKN_BRACE_OPEN:
		return parseObject(current, tokens)

	case TKN_BRACKET_OPEN:
		return parseArray(current, tokens)

	default:
		return nil, fmt.Errorf("invalid token type: %s", token.Type)
	}
}

func parseObject(current *int, tokens []Token) (ASTNode, error) {
	node := ObjectNode{
		Value: make(map[string]ASTNode),
	}

	*current++

	for *current < len(tokens) && tokens[*current].Type != TKN_BRACE_CLOSE {
		currToken := tokens[*current]

		if currToken.Type != TKN_STRING {
			return nil, fmt.Errorf("expected string key in object, got: %s", currToken.Type)
		}

		key := currToken.Value
		*current++

		if *current >= len(tokens) || tokens[*current].Type != TKN_COLON {
			return nil, fmt.Errorf("expected : in key value pair, got: %s", currToken.Type)
		}

		*current++

		value, err := parseValue(current, tokens)
		if err != nil {
			return nil, err
		}

		node.Value[key] = value

		if *current < len(tokens) && tokens[*current].Type == TKN_COMMA {
			*current++
		}
	}

	if *current >= len(tokens) || tokens[*current].Type != TKN_BRACE_CLOSE {
		return nil, fmt.Errorf("expected closing brace, got: %s", tokens[*current].Type)
	}

	*current++

	return node, nil
}

func parseArray(current *int, tokens []Token) (ASTNode, error) {
	node := ArrayNode{
		Value: make([]ASTNode, 0),
	}

	*current++

	for *current < len(tokens) && tokens[*current].Type != TKN_BRACKET_CLOSE {

		val, err := parseValue(current, tokens)
		if err != nil {
			return nil, err
		}

		node.Value = append(node.Value, val)

		if *current < len(tokens) && tokens[*current].Type == TKN_COMMA {
			*current++
		}
	}

	if *current >= len(tokens) || tokens[*current].Type != TKN_BRACKET_CLOSE {
		return nil, fmt.Errorf("expected closing bracket, got: %s", tokens[*current].Type)
	}

	*current++

	return node, nil
}
