package ioutils

import (
	"io"

	"github.com/aaron-vaz/golang-utils/pkg/errorutil"
)

// Close is helper method to close io.Closer resources
// It performs an error check and so can be used in deferred calls
func Close(resource io.Closer) {
	if resource != nil {
		errorutil.ErrCheck(resource.Close(), false)
	}
}
