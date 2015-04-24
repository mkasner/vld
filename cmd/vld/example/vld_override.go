package example

import "github.com/mkasner/vld"

// Validate override of Contract
// You can implement this in some other file
func (t *Contract) ValidateOverride(ctx vld.Context) vld.Context {
	return ctx
}
