package mysql

import (
	"github.com/j0urneyK/bob"
	"github.com/j0urneyK/bob/dialect/mysql/dialect"
)

func Insert(queryMods ...bob.Mod[*dialect.InsertQuery]) bob.BaseQuery[*dialect.InsertQuery] {
	q := &dialect.InsertQuery{}
	for _, mod := range queryMods {
		mod.Apply(q)
	}

	return bob.BaseQuery[*dialect.InsertQuery]{
		Expression: q,
		Dialect:    dialect.Dialect,
		QueryType:  bob.QueryTypeInsert,
	}
}
