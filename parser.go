// This file is generated - do not edit.

package definition

import (
	"github.com/elimity-com/abnf/operators"
)

// rule1 = "foo" bar baz
func Rule1(s []byte) operators.Alternatives {
	return operators.Concat(
		"rule1",
		operators.String("foo", "foo"),
		Bar,
		Baz,
	)(s)
}
