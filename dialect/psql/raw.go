package psql

import (
	"github.com/j0urneyK/bob"
	"github.com/j0urneyK/bob/dialect/psql/dialect"
	"github.com/j0urneyK/bob/expr"
)

func RawQuery(q string, args ...any) bob.BaseQuery[expr.Clause] {
	return expr.RawQuery(dialect.Dialect, q, args...)
}
