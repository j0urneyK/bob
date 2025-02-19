package sqlite_test

import (
	"testing"

	"github.com/j0urneyK/bob/dialect/sqlite"
	"github.com/j0urneyK/bob/dialect/sqlite/dm"
	testutils "github.com/j0urneyK/bob/test/utils"
)

func TestDelete(t *testing.T) {
	examples := testutils.Testcases{
		"simple": {
			Query: sqlite.Delete(
				dm.From("films"),
				dm.Where(sqlite.Quote("kind").EQ(sqlite.Arg("Drama"))),
			),
			ExpectedSQL:  `DELETE FROM films WHERE ("kind" = ?1)`,
			ExpectedArgs: []any{"Drama"},
		},
	}

	testutils.RunTests(t, examples, formatter)
}
