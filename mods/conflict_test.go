package mods

import (
	"github.com/j0urneyk/bob"
	"github.com/j0urneyk/bob/clause"
)

var _ bob.Mod[interface{ SetConflict(clause.Conflict) }] = Conflict[interface{ SetConflict(clause.Conflict) }](nil)
