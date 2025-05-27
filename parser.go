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
		return nil, fmt.Errorf("invalid token type: %s", token.Type)
	}
}

func parseObject(current *int, tokens []Token) (ASTNode, error) {
	node := ObjectNode{
		Value: make(map[string]ASTNode),
	}

	*current++

	currToken := tokens[*current]

	for currToken.Type != TKN_BRACE_CLOSE {
		if currToken.Type == TKN_STRING {
			key := currToken.Value

			*current++

			if tokens[*current].Type != TKN_COLON {
				return nil, fmt.Errorf("expected : in key value pair, got: %s", currToken.Type)
			}

			*current++

			value, err := parseValue(tokens[*current])
			if err != nil {
				return nil, err
			}

			node.Value[key] = value
		} else {
			return nil, fmt.Errorf("expected string key in object, got: %s", currToken.Type)
		}

		*current++

		if tokens[*current].Type == TKN_COMMA {
			*current++
		}
	}

	return node, nil
}
