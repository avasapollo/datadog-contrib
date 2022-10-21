package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

type Inter struct {
	Name    string
	Methods []*Method
}

type Method struct {
	Name   string
	Input  []*Param
	Output []*Param
}

func (m *Method) WithContext() bool {
	return strings.Contains(m.Name, "WithContext")
}

func (m *Method) GetParams() string {
	var inpt []string
	for _, p := range m.Input {
		inpt = append(inpt, fmt.Sprintf("%s %s", p.Name, p.Type))
	}
	return strings.Join(inpt, ",")
}

func (m *Method) GetResults() string {
	var inpt []string
	for _, p := range m.Output {
		inpt = append(inpt, p.Type)
	}
	return strings.Join(inpt, ",")
}

type Param struct {
	Name string
	Type string
}

func main() {
	//Create a FileSet to work with
	fset := token.NewFileSet()
	//Parse the file and create an AST
	file, err := parser.ParseFile(fset, "dynamo.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	list := iterInterfaces(file)
	inter := manipulate(list[0])
	f, err := os.Create("aws/v1/dynamodb/dynamodb.go")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("dynamodb.gotpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, inter)
	if err != nil {
		panic(err)
	}
}

func manipulate(n *namedInterface) *Inter {
	inter := &Inter{
		Name:    n.name.Name,
		Methods: []*Method{},
	}
	for _, method := range n.it.Methods.List {
		m := &Method{
			Name:   method.Names[0].Name,
			Input:  []*Param{},
			Output: []*Param{},
		}

		// inputs
		for _, pa := range method.Type.(*ast.FuncType).Params.List {
			tp := ""
			name := ""
			// pointer input
			if ptr, ok := pa.Type.(*ast.StarExpr); ok {
				name = "in"
				tp = "*" + ptr.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + ptr.X.(*ast.SelectorExpr).Sel.Name
			}
			// context
			if selector, ok := pa.Type.(*ast.SelectorExpr); ok {
				name = "ctx"
				tp = selector.X.(*ast.Ident).Name + "." + selector.Sel.Name
			}

			// opts
			if selector, ok := pa.Type.(*ast.Ellipsis); ok {
				name = "opts"
				tp = "..." + selector.Elt.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + selector.Elt.(*ast.SelectorExpr).Sel.Name
			}

			p := &Param{
				Name: name,
				Type: tp,
			}
			m.Input = append(m.Input, p)
		}

		// output
		for _, pa := range method.Type.(*ast.FuncType).Results.List {
			tp := ""
			// pointer input
			if ptr, ok := pa.Type.(*ast.StarExpr); ok {
				tp = "*" + ptr.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + ptr.X.(*ast.SelectorExpr).Sel.Name
			}
			// error
			if ind, ok := pa.Type.(*ast.Ident); ok {
				tp = ind.Name
			}

			p := &Param{
				Type: tp,
			}
			m.Output = append(m.Output, p)
		}

		inter.Methods = append(inter.Methods, m)
	}
	return inter
}

type namedInterface struct {
	name       *ast.Ident
	it         *ast.InterfaceType
	typeParams []*ast.Field
}

// Create an iterator over all interfaces in file.
func iterInterfaces(file *ast.File) []*namedInterface {
	var list []*namedInterface
	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			it, ok := ts.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			list = append(list, &namedInterface{ts.Name, it, getTypeSpecTypeParams(ts)})
		}
	}
	return list
}

func getTypeSpecTypeParams(ts *ast.TypeSpec) []*ast.Field {
	return nil
}
