package main

import (
	"fmt"
	"go/ast"
	"os"
	"strings"
)

type VldRules struct {
	List []VldRule
}

func (t *VldRules) byKey(key string) *VldRule {
	for _, r := range t.List {
		if r.Key == key {
			return &r
		}
	}
	return nil
}

type VldRule struct {
	Func    string
	Key     string
	Params  []string
	Package string
}

func (g *Generator) fetchRules() VldRules {
	rules := make([]VldRule, 0, 100)
	for _, file := range g.pkg.files {
		if file.file != nil {
			ast.Inspect(file.file, file.vldRuleDecl)
			rules = append(rules, file.vldRules...)
		}
	}
	return VldRules{List: rules}
}

// genDecl processes one declaration clause.
func (f *File) vldRuleDecl(node ast.Node) bool {
	fdecl, ok := node.(*ast.FuncDecl)
	if !ok {
		// We only care about func declarations.
		return true
	}
	// fmt.Printf("%+v\n", fdecl)
	ident := fdecl.Name
	fname := ident.Name

	doc := fdecl.Doc.Text()
	if !strings.Contains(doc, "vld") {
		return true
	}
	split := strings.Split(doc, ":")
	if len(split) < 1 {
		return true
	}
	tag := strings.TrimSpace(split[1])

	ftype := fdecl.Type
	paramsList := ftype.Params
	params := make([]string, 0)
	for _, p := range paramsList.List {
		if len(p.Names) == 0 {
			continue
		}
		for _, n := range p.Names {
			ident := n
			paramName := ident.Name
			params = append(params, paramName)
		}

	}
	f.vldRules = append(f.vldRules, VldRule{Func: fname, Key: tag, Params: params, Package: f.pkg.name})

	return false
}

// Helpers

func processRuleTag(tag string) []Rule {
	result := make([]Rule, 0)
	rules := strings.Split(tag, ",")
	for _, r := range rules {
		rv := strings.Split(r, "=")
		rule := Rule{}
		if len(rv) > 0 {
			rule.Key = rv[0]
		}
		if len(rv) == 2 {
			rule.Value = rv[1]
		}
		result = append(result, rule)
	}
	return result
}

func (g *Generator) fetchRuleRepo(rulePackages []string) VldRules {
	var vldg Generator
	gopath := ""
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "GOPATH" {
			gopath = pair[1]
		}
	}
	result := VldRules{}
	if gopath == "" {
		return result
	}
	for _, rp := range rulePackages {
		vldg.parsePackageDir(fmt.Sprintf("%s/src/%s", gopath, rp))
		rulesRepo := vldg.fetchRules()
		result.List = append(result.List, rulesRepo.List...)
	}
	// fmt.Printf("RulesRepo: %# v\n", pretty.Formatter(rulesRepo.List))
	return result
}
