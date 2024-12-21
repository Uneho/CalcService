package calculator

import (
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	if tokens == nil {
		return 0, ErrInvalidExpression
	}
	result, _, err := parseExpression(tokens, 0)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func tokenize(expression string) []string {
	var tokens []string
	for i := 0; i < len(expression); {
		switch {
		case unicode.IsDigit(rune(expression[i])) || expression[i] == '.':
			j := i
			for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
				j++
			}
			tokens = append(tokens, expression[i:j])
			i = j
		case expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/' || expression[i] == '(' || expression[i] == ')':
			tokens = append(tokens, string(expression[i]))
			i++
		case unicode.IsSpace(rune(expression[i])):
			i++
		default:
			return nil
		}
	}
	return tokens
}

func parseExpression(tokens []string, i int) (float64, int, error) {
	result, i, err := parseTerm(tokens, i)
	if err != nil {
		return 0, i, err
	}
	for i < len(tokens) {
		op := tokens[i]
		if op != "+" && op != "-" {
			break
		}
		i++
		next, newI, err := parseTerm(tokens, i)
		if err != nil {
			return 0, newI, err
		}
		i = newI
		if op == "+" {
			result += next
		} else {
			result -= next
		}
	}
	return result, i, nil
}

func parseTerm(tokens []string, i int) (float64, int, error) {
	result, i, err := parseFactor(tokens, i)
	if err != nil {
		return 0, i, err
	}
	for i < len(tokens) {
		op := tokens[i]
		if op != "*" && op != "/" {
			break
		}
		i++
		next, newI, err := parseFactor(tokens, i)
		if err != nil {
			return 0, newI, err
		}
		i = newI
		if op == "*" {
			result *= next
		} else {
			if next == 0 {
				return 0, i, ErrDivisionByZero
			}
			result /= next
		}
	}
	return result, i, nil
}

func parseFactor(tokens []string, i int) (float64, int, error) {
	if i >= len(tokens) {
		return 0, i, ErrUnexpectedEndOfExpr
	}
	if tokens[i] == "(" {
		i++
		result, newI, err := parseExpression(tokens, i)
		if err != nil {
			return 0, newI, err
		}
		if newI >= len(tokens) || tokens[newI] != ")" {
			return 0, newI, ErrMismatchedParentheses
		}
		return result, newI + 1, nil
	}
	value, err := strconv.ParseFloat(tokens[i], 64)
	if err != nil {
		return 0, i, ErrInvalidNumber
	}
	return value, i + 1, nil
}
