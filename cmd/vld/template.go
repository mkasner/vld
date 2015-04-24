package main

const (
	vldFn = `// Validates {{.TypeName}}
func (t *{{.TypeName}}) Validate(ctx vld.Context) vld.Context {
	validator := vld.New().
	{{$len := len .Validations}}
	{{$typeName := .TypeName}}
	{{range $i, $v := .Validations}}Add({{$v.Package}}{{if $v.Package}}.{{end}}{{$v.Func}}({{if $v.Value}}{{$v.Value}},{{end}} t.{{$v.Field}}, "{{$typeName}}.{{$v.Field}}")){{if NotLast $len $i}}.{{end}}
	{{end}}

	ctx.Err = validator.Validate()
	
	ctx = t.ValidateOverride(ctx)
	return ctx
}
`

	vldOverrirdeFn = `// Validate override of {{.TypeName}}
// You can implement this in some other file
func (t *{{.TypeName}}) {{OverrideFn}}(ctx vld.Context)  vld.Context {
	return ctx
}
`
)
