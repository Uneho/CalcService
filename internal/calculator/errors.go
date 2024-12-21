package calculator

import "errors"

var (
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrDivisionByZero        = errors.New("division by zero")
	ErrUnexpectedEndOfExpr   = errors.New("unexpected end of expression")
	ErrMismatchedParentheses = errors.New("mismatched parentheses")
	ErrInvalidNumber         = errors.New("invalid number")
)

func GetErrorMessage(err error) string {
	switch err {
	case ErrInvalidExpression:
		return "Expression is not valid"
	case ErrDivisionByZero:
		return "Division by zero is not allowed"
	case ErrUnexpectedEndOfExpr:
		return "Unexpected end of expression"
	case ErrMismatchedParentheses:
		return "Mismatched parentheses in expression"
	case ErrInvalidNumber:
		return "Invalid number in expression"
	default:
		return "Internal server error"
	}
}
