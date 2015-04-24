// Package for declaring and executing validations which can be techninal validations,
// bussiness validations, struct validations
package vld

// Passes data between validations.
// Used for error passing at the moment.
type Context struct {
	Err error
}

// Container for all validation errors.
// Used when ValidateAll function call
type VldErrors struct {
	Err []error
}

// Every type which wants to act as validation must implement this interface
type Validation interface {
	Validate(Context) Context
	ValidateOverride(Context) Context
}

// Type that holds validation func which implements Validation interface
type ValidationFunc func() error

func (f ValidationFunc) Validate(ctx Context) Context {
	ctx.Err = f()
	ctx = f.ValidateOverride(ctx)
	return ctx
}

func (f ValidationFunc) ValidateOverride(ctx Context) Context {
	return ctx
}

// Default struct which you can embed into your struct that implements ValidateOverride method.
// In that way you only must implement Validate method
type DefaultValidation struct {
}

func (f DefaultValidation) ValidateOverride(ctx Context) Context {
	return ctx
}

// Holds and runs validations
type Validator struct {
	validations []Validation
}

// Validate and exit at first error
func (t *Validator) Validate() error {
	ctx := Context{}
	for _, v := range t.validations {
		if ctx.Err != nil {
			continue
		}
		ctx = v.Validate(ctx)
	}
	return ctx.Err
}

// Validates all validations and collects errors
func (t *Validator) ValidateAll() VldErrors {
	ctx := Context{}
	errs := VldErrors{Err: make([]error, 0)}

	for _, v := range t.validations {
		if ctx.Err != nil {
			errs.Err = append(errs.Err, ctx.Err)
		}
		ctx = v.Validate(ctx)
	}
	return errs
}

// Adds validation to validator
func (t *Validator) Add(v Validation) *Validator {
	t.validations = append(t.validations, v)
	return t
}

// Add regular func to validator
// It will act as validation on validator
func (t *Validator) AddFunc(f func() error) *Validator {
	// vf := func() error {
	// 	return f()
	// }
	t.validations = append(t.validations, ValidationFunc(f))
	return t
}

// Returns validations on this validator
func (t *Validator) Validations(v Validation) []Validation {
	return t.validations
}

// Creates new Validator on which you can chain validations
func New() *Validator {
	return &Validator{validations: make([]Validation, 0)}
}
