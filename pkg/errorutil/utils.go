package errorutil

import (
	"log"
	"os"
	"runtime/debug"
)

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
