package data

import (
	"github.com/davecgh/go-spew/spew" // nolint: depguard
)

// Dumper will always disable Error() and Stringer() methods which is otherwize
// default in spew.
func Dumper(i interface{}) {
	cfg := spew.NewDefaultConfig()
	cfg.DisableMethods = true
	cfg.Dump(i)
}
