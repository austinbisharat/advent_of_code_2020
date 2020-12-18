package math

import (
	"fmt"
	"strconv"
	"strings"
)

func EvaluateExpression(exp string) (int, int, error) {
	tokens := tokenizeExpression(exp)
	ast1, _, err := parseTokensLeftRight(tokens)
	if err != nil {
		return 0, 0, err
	}
	ast2, _, err := parseTokensPrecedence(tokens)
	if err != nil {
		return 0, 0, err
	}

	value1, err := evaluateNode(ast1)
	if err != nil {
		return 0, 0, err
	}
	value2, err := evaluateNode(ast2)
	if err != nil {
		return 0, 0, err
	}
	return value1, value2, nil
}

type token string

const (
	openParen  token = "("
	closeParen token = ")"
	plus       token = "+"
	times      token = "*"
)

func tokenizeExpression(exp string) []token {
	stringParts := strings.Split(exp, " ")
	tokens := []token{openParen}
	for _, t := range stringParts {
		for len(t) > 1 && t[0] == '(' {
			tokens = append(tokens, openParen)
			t = t[1:]
		}
		countEnd := 0
		for len(t) > 1 && t[len(t)-1] == ')' {
			countEnd++
			t = t[:len(t)-1]
		}

		if len(t) > 0 {
			tokens = append(tokens, token(t))
		}

		for i := 0; i < countEnd; i++ {
			tokens = append(tokens, closeParen)
		}

	}
	tokens = append(tokens, closeParen)
	return tokens
}

type astNodeType int

const (
	astNodeTypeAdd astNodeType = iota + 1
	astNodeTypeMultiply
	astNodeTypeValue
)

type astNode struct {
	nodeType astNodeType
	value    int
	left     *astNode
	right    *astNode
}

func evaluateNode(node *astNode) (int, error) {
	if node == nil {
		return 0, fmt.Errorf("cannot evaluate nil node")
	}
	switch node.nodeType {
	case astNodeTypeValue:
		return node.value, nil
	case astNodeTypeAdd:
		l, err := evaluateNode(node.left)
		if err != nil {
			return 0, err
		}
		r, err := evaluateNode(node.right)
		if err != nil {
			return 0, err
		}
		return r + l, nil
	case astNodeTypeMultiply:
		l, err := evaluateNode(node.left)
		if err != nil {
			return 0, err
		}
		r, err := evaluateNode(node.right)
		if err != nil {
			return 0, err
		}
		return r * l, nil
	default:
		return 0, fmt.Errorf("unknown node type %d", node.nodeType)
	}
}

func parseTokensLeftRight(tokens []token) (*astNode, []token, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("cannot parse 0 tokens")
	}

	firstToken := tokens[0]
	if firstToken == closeParen || firstToken == plus || firstToken == times {
		return nil, nil, fmt.Errorf("first token of expression must be operator or open paren: %s", firstToken)
	}

	if firstToken == "(" {
		var node *astNode
		var err error

		node, tokens, err = parseTokensLeftRight(tokens[1:])
		if err != nil {
			return nil, nil, err
		}

		for len(tokens) > 0 && tokens[0] != closeParen {
			parentNode := &astNode{}
			if tokens[0] == plus {
				parentNode.nodeType = astNodeTypeAdd
			} else if tokens[0] == times {
				parentNode.nodeType = astNodeTypeMultiply
			} else {
				return nil, nil, fmt.Errorf("second token in expression must be * or +, but got %s", tokens[0])
			}
			parentNode.left = node
			parentNode.right, tokens, err = parseTokensLeftRight(tokens[1:])
			if err != nil {
				return nil, nil, err
			}
			node = parentNode
		}

		if len(tokens) == 0 {
			return nil, nil, fmt.Errorf("expected closing paren")
		}
		return node, tokens[1:], nil
	}

	value, err := strconv.Atoi(string(firstToken))
	node := &astNode{
		nodeType: astNodeTypeValue,
		value:    value,
		left:     nil,
		right:    nil,
	}
	return node, tokens[1:], err
}

func parseTokensPrecedence(tokens []token) (*astNode, []token, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("cannot parse 0 tokens")
	}

	firstToken := tokens[0]
	if firstToken == closeParen || firstToken == plus || firstToken == times {
		return nil, nil, fmt.Errorf("first token of expression must be operator or open paren: %s", firstToken)
	}

	if firstToken == "(" {
		var node *astNode
		var err error

		node, tokens, err = parseTokensPrecedence(tokens[1:])
		if err != nil {
			return nil, nil, err
		}

		prevNodeType := astNodeTypeValue
		for len(tokens) > 0 && tokens[0] != closeParen {
			newNode := &astNode{}
			if tokens[0] == plus {
				newNode.nodeType = astNodeTypeAdd
			} else if tokens[0] == times {
				newNode.nodeType = astNodeTypeMultiply
			} else {
				return nil, nil, fmt.Errorf("second token in expression must be * or +, but got %s", tokens[0])
			}

			if prevNodeType == astNodeTypeMultiply && newNode.nodeType == astNodeTypeAdd {
				newNode.left = node.right
				newNode.right, tokens, err = parseTokensPrecedence(tokens[1:])
				node.right = newNode
			} else {
				newNode.left = node
				newNode.right, tokens, err = parseTokensPrecedence(tokens[1:])
				if err != nil {
					return nil, nil, err
				}
				node = newNode
			}

			prevNodeType = node.nodeType
		}

		if len(tokens) == 0 {
			return nil, nil, fmt.Errorf("expected closing paren")
		}
		return node, tokens[1:], nil
	}

	value, err := strconv.Atoi(string(firstToken))
	node := &astNode{
		nodeType: astNodeTypeValue,
		value:    value,
		left:     nil,
		right:    nil,
	}
	return node, tokens[1:], err
}
