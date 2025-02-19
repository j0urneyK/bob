package dialect

import (
	"context"
	"strings"

	"github.com/j0urneyK/bob"
	"github.com/j0urneyK/bob/expr"
)

type Expression struct {
	expr.Chain[Expression, Expression]
}

func (Expression) New(exp bob.Expression) Expression {
	var b Expression
	b.Base = exp
	return b
}

// Implements fmt.Stringer()
func (x Expression) String() string {
	w := strings.Builder{}
	x.WriteSQL(context.Background(), &w, Dialect, 1) //nolint:errcheck
	return w.String()
}
