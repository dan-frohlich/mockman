package main

import (
	"os"
	"text/template"
	"unicode"
	"unicode/utf8"
)

func stubInterface(packageName string, info InterfaceInfo) string {
	data := StubInfo{
		Package:   packageName,
		Interface: info,
	}

	funcMap := template.FuncMap{
		"lcase": lCaseGoSymbol,
	}

	t := template.Must(template.New("userTemplate").Funcs(funcMap).Parse(tpl))
	t.Execute(os.Stdout, data)
	return ""
}

type StubInfo struct {
	Package   string
	Interface InterfaceInfo
}

var tpl = `package {{.Package}}

type {{.Interface.Name}}Stub struct {
{{range .Interface.Methods}}
  {{lcase .Name}} func({{range $index, $param := .Parameters}}{{if $index}}, {{end}}{{.Name}}{{.Type}}{{end}}) {{if eq (len .Results) 0}}{{else if eq (len .Results) 1}}{{index .Results 0}}{{else}}({{range $index, $result := .Results}}{{if $index}}, {{end}}{{.}}{{end}}){{end}}
{{end -}}
}

{{ $ifName := .Interface.Name }}

{{range $index, $func := .Interface.Methods}}
func (stub *{{$ifName}}Stub) {{$func }}{
	if stub.{{lcase $func.Name}} != nil {
		return stub.{{lcase $func.Name}}({{range $i, $param := $func.Parameters}}{{if $i}}, {{end}}{{.Name}}{{end}})
	}
	panic("Method {{$func.Name}} not implemented in {{$ifName}}Stub")
}
{{end}}
`

/*
{{range .Interface.Methods}}
func (s *{{.Interface.Name}}Stub) {{range .Interface.Methods}}{{.String}} {
	if s.{{lcase .Name}} != nil {
		return s.{{lcase .Name}}({{range $index, $param := .Parameters}}{{if $index}}, {{end}}{{.Name}}{{end}}{{end}})
	}
	panic("Method {{.Name}} not implemented in {{.Interface.Name}}Stub")
}
{{end}}

*/

func lCaseGoSymbol(goSymbol string) string {
	if goSymbol == "" {
		return goSymbol
	}
	r, i := utf8.DecodeRuneInString(goSymbol)
	lc := unicode.ToLower(r)
	return string(lc) + goSymbol[i:]
}
