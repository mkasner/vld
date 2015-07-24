package main

import (
	"strings"
	"testing"
)

type testStruct struct {
	tag   StructTag
	key   string
	value string
}

var (
	testData = []testStruct{
		testStruct{tag: `db:"name" vld:"req,maxlen=100"`, key: "db", value: "name"},
		testStruct{tag: `db:"name" vld:"req,maxlen=100"`, key: "vld", value: "req"},
	}
)

func TestStructTagGet(t *testing.T) {
	for _, td := range testData {
		keyData := td.tag.Get(td.key)
		if !strings.Contains(keyData, td.value) {
			t.Errorf("tag: %s  key: %s expected: %s  result: %s", td.tag, td.key, td.value, keyData)
			t.Fail()
		}

	}
}
