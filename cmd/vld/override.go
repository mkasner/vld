package main

import (
	"go/ast"
	"strings"
)

const (
	FN_OVERRIDE = "ValidateOverride"
)

func (g *Generator) fetchOverride(typeName string) bool {
	var result bool
	for _, file := range g.pkg.files {
		if file.file != nil {
			file.typeName = typeName
			file.overrideExist = false
			ast.Inspect(file.file, file.vldOverrideDecl)
			result = file.overrideExist
		}
	}
	return result
}

// genDecl processes one declaration clause.
func (f *File) vldOverrideDecl(node ast.Node) bool {
	fdecl, ok := node.(*ast.FuncDecl)
	if !ok {
		// We only care about func declarations.
		return true
	}
	// fmt.Printf("%+v\n", fdecl)

	ident := fdecl.Name
	fname := ident.Name

	frecv := fdecl.Recv
	for _, p := range frecv.List {
		ptype := p.Type.(*ast.StarExpr)

		ident := ptype.X.(*ast.Ident)
		typ := ident.Name

		// Ignore type we are not looking for
		if typ != f.typeName {
			return true
		}

		// Ignore func we are not looking for
		if strings.TrimSpace(fname) != FN_OVERRIDE {
			return true
		}
		f.overrideExist = validateOverride(f, fname, typ)
	}

	return false
}

func validateOverride(f *File, fname, typ string) bool {
	if typ != f.typeName {
		return false
	}

	if strings.TrimSpace(fname) == FN_OVERRIDE {
		return true
	}
	return false
}
