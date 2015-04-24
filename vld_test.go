package vld

import (
	"fmt"
	"testing"
)

type stringValidation struct {
	DefaultValidation
	payload string
}

func (t *stringValidation) Validate(ctx Context) Context {
	if len(t.payload) > 4 {
		ctx.Err = fmt.Errorf("Too long: %s", t.payload)
	}
	return ctx
}

func TestBuildValidator(t *testing.T) {
	sv := &stringValidation{payload: "4211111132"}
	validator := New()
	payload := "helo"
	valueint := int64(1000)
	empty := 0

	validator = validator.
		AddFunc(func() error {
		if len(payload) > 4 {
			return fmt.Errorf("Too long: %s", payload)
		}
		return nil
	}).
		Add(Required(empty, "empty")).
		Add(MaxInt64(1000, valueint, "")).
		Add(MaxLength(4, payload, "")).
		Add(sv).
		Add(&stringValidation{payload: "42323312"})
	err := validator.Validate()
	if err != nil {
		t.Log(err.Error())
	}
}
