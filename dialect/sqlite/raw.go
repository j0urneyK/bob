package sqlite

import (
	"github.com/j0urneyk/bob"
	"github.com/j0urneyk/bob/dialect/sqlite/dialect"
	"github.com/j0urneyk/bob/expr"
)

func RawQuery(q string, args ...any) bob.BaseQuery[expr.Clause] {
	return expr.RawQuery(dialect.Dialect, q, args...)
}
