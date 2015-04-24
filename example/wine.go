//go:generate vld -type=Wine
package example

import "github.com/mkasner/vld"

type Wine struct {
	Kind string `vld:"req,maxlen=10"`
	Year int    `vld:"req,maxint=3000,minint=1900"`
}

func isValid(w Wine) bool {
	if err := vld.New().Add(&w).Validate(); err != nil {
		return true
	}
	return false
}
