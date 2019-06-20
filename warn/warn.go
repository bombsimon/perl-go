package warn

import (
	"fmt"
	"os"
	"runtime"
)

// Here will warn at the current line.
func Here(m ...string) {
	_, file, no, _ := runtime.Caller(1)

	message := "Warning: something's wrong"

	if len(m) > 0 {
		message = m[0]
	}

	fmt.Fprintf(os.Stderr, "%s at %s line %d.\n", message, file, no)
}
