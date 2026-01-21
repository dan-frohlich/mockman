package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func matchFilter(name, filter string) bool {
	return strings.Contains(strings.ToLower(name), strings.ToLower(filter))
}

func makeMethodInfo(methodName *ast.Ident, method *ast.Field) (MethodInfo, bool) {

	mi := MethodInfo{
		Name: methodName.Name,
	}
	funcType, ok := method.Type.(*ast.FuncType)
	if !ok {
		return mi, false
	}

	if funcType.Params != nil {
		for _, param := range funcType.Params.List {
			paramType := exprToString(param.Type)
			for _, paramName := range param.Names {
				mi.Parameters = append(mi.Parameters, ParamInfo{
					Name: paramName.Name,
					Type: paramType,
				})
			}
			if len(param.Names) == 0 {
				mi.Parameters = append(mi.Parameters, ParamInfo{
					Name: "",
					Type: paramType,
				})
			}
		}
	}
	if funcType.Results != nil {
		for _, result := range funcType.Results.List {
			resultType := exprToString(result.Type)
			mi.Results = append(mi.Results, resultType)
		}
	}
	return mi, true
}

func findInterfaces(path, filter string) ([]InterfaceInfo, error) {
	var interfaces []InterfaceInfo

	files, err := filepath.Glob(filepath.Join(path, "*.go"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, content, 0)
		if err != nil {
			continue
		}

		for _, decl := range f.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec := spec.(*ast.TypeSpec)
				if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
					if filter == "" || matchFilter(typeSpec.Name.Name, filter) {
						ii := InterfaceInfo{
							Name:   typeSpec.Name.Name,
							Source: interfaceType,
						}
						for _, method := range interfaceType.Methods.List {
							for _, methodName := range method.Names {
								mi, ok := makeMethodInfo(methodName, method)
								if ok {
									ii.Methods = append(ii.Methods, mi)
								}
							}
						}
						interfaces = append(interfaces, ii)
					}
				}
			}
		}
	}

	return interfaces, nil
}

// exprToString converts an ast.Expr to its string representation.
func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return exprToString(e.X) + "." + e.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(e.X)
	case *ast.ArrayType:
		return "[]" + exprToString(e.Elt)
	case *ast.Ellipsis:
		return "..." + exprToString(e.Elt)
	case *ast.MapType:
		return "map[" + exprToString(e.Key) + "]" + exprToString(e.Value)
	case *ast.FuncType:
		return "func"
	default:
		return ""
	}
}
