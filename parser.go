package main

import (
	"errors"
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

	return nil, nil
}

func parseValue(token Token) (ASTNode, error) {
	switch token.Type {
	case TKN_STRING:
		return StringNode{Value: token.Value}, nil

	case TKN_NUMBER:
		num, _ := strconv.ParseFloat(token.Value, 64)
		return NumberNode{Value: num}, nil

	case TKN_TRUE:
		return BooleanNode{Value: true}, nil
	case TKN_FALSE:
		return BooleanNode{Value: false}, nil
	case TKN_NULL:
		return NullNode{}, nil
	default:
		return nil, errors.New("invalid token type: " + token.Type)
	}
}
