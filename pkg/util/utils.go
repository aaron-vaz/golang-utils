package util

import (
	"io"
	"log"
	"os"
	"runtime/debug"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Close is helper method to close io.Closer resources
// It performs an error check and so can be used in deferred calls
func Close(resource io.Closer) {
	if resource != nil {
		ErrCheck(resource.Close(), false)
	}
}

// ErrCheck is a helper method to check is a error object is a valid error
// if the error is a valid it will log it
func ErrCheck(err error, exit bool) {
	if err != nil {
		log.Println(err, "\n", string(debug.Stack()))

		if exit {
			os.Exit(1)
		}
	}
}
