package main

import (
	"fmt"
	"testing"
)

func TestExtractFields(t *testing.T) {
	var g Generator
	// names := []string{"sqlscan_test.go"}
	g.parsePackageDir("./example")
	rulePackages := []string{"github.com/mkasner/vld"}
	g.generate("Transaction", rulePackages)
	fmt.Println(g.buf.String())
}

func TestOverrideFuncExist(t *testing.T) {
	var g Generator
	// names := []string{"sqlscan_test.go"}
	g.parsePackageDir("./example")
	overrideExist := g.fetchOverride("Contract")
	if !overrideExist {
		t.Fail()
	}
}
