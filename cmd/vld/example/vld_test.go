package example

import (
	"testing"

	"github.com/mkasner/vld"
)

func TestBuildValidator(t *testing.T) {

	status := 1

	contract := &Contract{
		Price: 22,
	}

	invoice := &Invoice{
		Amount:   1.1,
		Email:    "mislav@@aduro.hr",
		Status:   &status,
		Contract: contract,
		MSISDN:   "38595430929",
	}

	transaction := &Transaction{
		Hello: "hellow",
		World: "world",
	}

	transaction2 := &Transaction{
		Hello: "helladads",
		World: "worldsdadasddada",
	}

	// Create validator
	validator := vld.New().
		Add(invoice).
		Add(contract).
		Add(transaction).
		Add(transaction2)

	// Run validator, exit at first error
	err := validator.Validate()
	if err != nil {
		t.Log(err.Error())
	}
}

func TestValidateAll(t *testing.T) {

	contract := &Contract{
		Price: 22,
	}
	transaction := &Transaction{
		Hello: "hellow",
		World: "world",
	}

	transaction2 := &Transaction{
		Hello: "helladads",
		World: "worldsdadasddada",
	}

	// Create validator
	validator := vld.New().
		Add(contract).
		Add(transaction).
		Add(transaction2)

	// Validate all validations
	errs := validator.ValidateAll()
	if len(errs.Err) > 0 {
		t.Logf("%+v\n", errs)
	}
}
