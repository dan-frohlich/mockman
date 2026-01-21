package main

import (
	"go/ast"
	"strings"
)

type interfaceInfo struct {
	name    string
	methods []methodInfo
	source  *ast.InterfaceType
}

func (ii interfaceInfo) Describe() string {
	var methods []string
	for _, m := range ii.methods {
		methods = append(methods, m.String())
	}
	return ii.name + " [" + strings.Join(methods, ", ") + "]"
}

type methodInfo struct {
	name       string
	parameters []paramInfo
	results    []string
}

func (mi methodInfo) String() string {
	var params []string
	for _, p := range mi.parameters {
		params = append(params, p.String())
	}
	s := mi.name + "(" + strings.Join(params, ", ") + ")"
	switch len(mi.results) {
	case 0:
		return s
	case 1:
		s += " " + mi.results[0]
	default:
		s += " (" + strings.Join(mi.results, ", ") + ")"
	}

	return s
}

type paramInfo struct {
	name string
	typ  string
}

func (pi paramInfo) String() string {
	return pi.name + " " + pi.typ
}
