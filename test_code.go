package main

import (
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

func mockInterface(packageName string, info InterfaceInfo) (string, error) {
	data := MockInfo{
		Package:   packageName,
		Interface: info,
	}

	funcMap := template.FuncMap{
		"lcase": lCaseGoSymbol,
	}

	t := template.Must(template.New("userTemplate").Funcs(funcMap).Parse(tpl))
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type MockInfo struct {
	Package   string
	Interface InterfaceInfo
}

var tpl = `package {{.Package}}

type {{.Interface.Name}}Mock struct {
{{range .Interface.Methods}}
  {{lcase .Name}} func({{range $index, $param := .Parameters}}{{if $index}}, {{end}}{{.Name}} {{.Type}}{{end}}) {{if eq (len .Results) 0}}{{else if eq (len .Results) 1}}{{index .Results 0}}{{else}}({{range $index, $result := .Results}}{{if $index}}, {{end}}{{.}}{{end}}){{end -}}

{{end -}}
}

{{ $ifName := .Interface.Name }}

{{range $index, $func := .Interface.Methods}}
func (mock *{{$ifName}}Mock) {{$func.Name}}({{range $i, $param := $func.Parameters}}{{if $i}}, {{end}}{{.String}}{{end}}) {{if eq (len $func.Results) 0}}{{else if eq (len $func.Results) 1}}{{index $func.Results 0}}{{else}}({{range $i, $result := $func.Results}}{{if $i}}, {{end}}{{.}}{{end}}){{end}} {
	if mock.{{lcase $func.Name}} != nil {
		return mock.{{lcase $func.Name}}({{range $i, $param := $func.Parameters}}{{if $i}}, {{end}}{{.Name}}{{end}})
	}
	panic("Method {{$func.Name}} not implemented in {{$ifName}}Mock")
}
{{end}}
`

/*
{{range .Interface.Methods}}
func (s *{{.Interface.Name}}Mock) {{range .Interface.Methods}}{{.String}} {
	if s.{{lcase .Name}} != nil {
		return s.{{lcase .Name}}({{range $index, $param := .Parameters}}{{if $index}}, {{end}}{{.Name}}{{end}}{{end}})
	}
	panic("Method {{.Name}} not implemented in {{.Interface.Name}}Mock")
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
