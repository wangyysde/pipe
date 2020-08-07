// +build js nacl plan9

package bzhylog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
