package fm

import (
	"github.com/j0urneyk/bob"
	"github.com/j0urneyk/bob/clause"
	"github.com/j0urneyk/bob/dialect/mysql/dialect"
	"github.com/j0urneyk/bob/mods"
)

func Distinct() bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.Distinct = true
	})
}

func OrderBy(e any) dialect.OrderBy[*dialect.Function] {
	return dialect.OrderBy[*dialect.Function](func() clause.OrderDef {
		return clause.OrderDef{
			Expression: e,
		}
	})
}

func Filter(e ...any) bob.Mod[*dialect.Function] {
	return bob.ModFunc[*dialect.Function](func(f *dialect.Function) {
		f.Filter = append(f.Filter, e...)
	})
}

func Over(winMods ...bob.Mod[*clause.Window]) bob.Mod[*dialect.Function] {
	w := clause.Window{}
	for _, mod := range winMods {
		mod.Apply(&w)
	}

	return mods.Window[*dialect.Function](w)
}
