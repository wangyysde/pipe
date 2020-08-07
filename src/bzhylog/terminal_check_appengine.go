// +build appengine

package bzhylog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
