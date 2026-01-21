package main

import (
	"go/ast"
	"strings"
)

type InterfaceInfo struct {
	Name    string
	Methods []MethodInfo
	Source  *ast.InterfaceType
}

func (ii InterfaceInfo) Describe() string {
	var methods []string
	for _, m := range ii.Methods {
		methods = append(methods, m.String())
	}
	return ii.Name + " [" + strings.Join(methods, ", ") + "]"
}

type MethodInfo struct {
	Name       string
	Parameters []ParamInfo
	Results    []string
}

func (mi MethodInfo) String() string {
	var params []string
	for _, p := range mi.Parameters {
		params = append(params, p.String())
	}
	s := mi.Name + "(" + strings.Join(params, ", ") + ")"
	switch len(mi.Results) {
	case 0:
		return s
	case 1:
		s += " " + mi.Results[0]
	default:
		s += " (" + strings.Join(mi.Results, ", ") + ")"
	}

	return s
}

type ParamInfo struct {
	Name string
	Type string
}

func (pi ParamInfo) String() string {
	return pi.Name + " " + pi.Type
}
