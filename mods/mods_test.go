package mods

import (
	"github.com/j0urneyk/bob"
	"github.com/j0urneyk/bob/clause"
)

var (
	_ bob.Mod[any]                                 = QueryMods[any](nil)
	_ bob.Mod[interface{ AppendWith(clause.CTE) }] = With[interface{ AppendWith(clause.CTE) }]{}
)
