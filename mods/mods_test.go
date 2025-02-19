package mods

import (
	"github.com/j0urneyK/bob"
	"github.com/j0urneyK/bob/clause"
)

var (
	_ bob.Mod[any]                                 = QueryMods[any](nil)
	_ bob.Mod[interface{ AppendWith(clause.CTE) }] = With[interface{ AppendWith(clause.CTE) }]{}
)
