package psql

import (
	"github.com/j0urneyk/bob"
	"github.com/j0urneyk/bob/dialect/psql/dialect"
	"github.com/j0urneyk/bob/expr"
)

func RawQuery(q string, args ...any) bob.BaseQuery[expr.Clause] {
	return expr.RawQuery(dialect.Dialect, q, args...)
}
