package mods

import (
	"github.com/j0urneyK/bob"
	"github.com/j0urneyK/bob/clause"
)

var _ bob.Mod[interface{ SetConflict(clause.Conflict) }] = Conflict[interface{ SetConflict(clause.Conflict) }](nil)
