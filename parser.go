package main

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
