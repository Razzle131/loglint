package logcheck

import (
	"go/ast"
)

func withTypeCheck[T any](src ast.Expr, next func(T) []error) []error {
	checked, ok := src.(T)
	if !ok {
		return []error{}
	}

	return next(checked)
}
