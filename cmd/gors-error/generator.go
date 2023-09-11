package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/constant"
	"go/format"
	"go/token"
	"go/types"
	"html/template"
	"log"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"
)

type errorInfo struct {
	Name     string // 错误名
	Desc     string // 错误描述
	Code     int    // 错误业务码
	HTTPCode string // 错误网络码
}

type errorWrapper struct {
	Errors []*errorInfo
}

var errTemp = `
var (
{{ range .Errors }}

Err{{.Name}} = gors.Error{
	StatusCode: {{.HTTPCode}},
	Code: {{.Name}},
	Message: "{{.Desc}}",
}.Froze()

{{- end }}
)
`

var errCodeDocPrefix = `# 错误码

！！系统错误码列表，由 {{.}}gen_error -type=int -doc{{.}} 命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中存在 {{.}}code{{.}} 字段，则表示调用 API 接口失败。例如：

{{.}}{{.}}{{.}}json
{
  "code": 100101,
  "message": "Database error"
}
{{.}}{{.}}{{.}}

上述返回中 {{.}}code{{.}} 表示错误码，{{.}}message{{.}} 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
`

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.

	trimPrefix string
}

// Printf like fmt.Printf, but add the string to g.buf.
func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the constant type.
	values   []Value // Accumulator for constant values of that type.

	trimPrefix string
}

// Package defines options for package.
type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(srcLocaltion string, patterns []string, tags []string) {
	cfg := &packages.Config{
		Dir: srcLocaltion,
		// nolint: staticcheck
		Mode: packages.LoadSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0])
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file:       file,
			pkg:        g.pkg,
			trimPrefix: g.trimPrefix,
		}
	}
}

// generateDocs produces error code markdown document for the named type.
func (g *Generator) generateDocs(typeName string) {
	values := make([]Value, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		file.values = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			values = append(values, file.values...)
		}
	}

	if len(values) == 0 {
		log.Printf("no values defined for type %s\n", typeName)
		return
	}

	tmpl, _ := template.New("doc").Parse(errCodeDocPrefix)
	var buf bytes.Buffer
	_ = tmpl.Execute(&buf, "`")

	// Generate code that will fail if the constants change value.
	g.Printf(buf.String())
	for _, v := range values {
		code, description := v.ParseComment()
		// g.Printf("\tregister(%s, %s, \"%s\")\n", v.originalName, code, description)
		g.Printf("| %s | %d | %s | %s |\n", v.originalName, v.value, code, description)
	}
	g.Printf("\n")
}

// generateErrs produces error info to make error functions.
func (g *Generator) generate(typeName string) {
	values := make([]Value, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		file.values = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			values = append(values, file.values...)
		}
	}

	if len(values) == 0 {
		log.Fatalf("no values defined for type %s", typeName)
	}

	var ew errorWrapper
	for _, v := range values {
		code, description := v.ParseComment()
		err := &errorInfo{
			Name:     v.originalName,
			Desc:     description,
			Code:     int(v.value),
			HTTPCode: code,
		}
		ew.Errors = append(ew.Errors, err)
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errTemp)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, &ew); err != nil {
		panic(err)
	}

	// Generate code that will fail if the constants change value.
	g.Printf(buf.String())
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")

		return g.buf.Bytes()
	}

	return src
}

// Value represents a declared constant.
type Value struct {
	comment      string
	originalName string // The name of the constant.
	name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value  uint64 // Will be converted to int64 when needed.
	signed bool   // Whether the constant is a signed type.
	str    string // The string representation given by the "go/constant" package.
}

func (v *Value) String() string {
	return v.str
}

// ParseComment parse comment to http code and error code description.
func (v *Value) ParseComment() (string, string) {
	reg := regexp.MustCompile(`\w\s*-\s*(\d{3})\s*:\s*([\w\W]*)\s*\.\n*`)
	if !reg.MatchString(v.comment) {
		log.Printf("constant '%s' have wrong comment format, register with 500 as default", v.originalName)

		return "500", "Internal server error"
	}

	groups := reg.FindStringSubmatch(v.comment)
	if len(groups) != 3 {
		return "500", "Internal server error"
	}

	return groups[1], groups[2]
}

// nolint: gocognit
// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		vspec, _ := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
		if vspec.Type == nil && len(vspec.Values) > 0 {
			// "X = 1". With no type but a value. If the constant is untyped,
			// skip this vspec and reset the remembered type.
			typ = ""

			// If this is a simple type conversion, remember the type.
			// We don't mind if this is actually a call; a qualified call won't
			// be matched (that will be SelectorExpr, not Ident), and only unusual
			// situations will result in a function call that appears to be
			// a type conversion.
			ce, ok := vspec.Values[0].(*ast.CallExpr)
			if !ok {
				continue
			}
			id, ok := ce.Fun.(*ast.Ident)
			if !ok {
				continue
			}
			typ = id.Name
		}
		if vspec.Type != nil {
			// "X T". We have a type. Remember it.
			ident, ok := vspec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			typ = ident.Name
		}
		if typ != f.typeName {
			// This is not the type we're looking for.
			continue
		}
		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.
		// Grab their names and actual values and store them in f.values.
		for _, name := range vspec.Names {
			if name.Name == "_" {
				continue
			}
			// This dance lets the type checker find the values for us. It's a
			// bit tricky: look up the object declared by the name, find its
			// types.Const, and extract its value.
			obj, ok := f.pkg.defs[name]
			if !ok {
				log.Fatalf("no value for constant %s", name)
			}
			info := obj.Type().Underlying().(*types.Basic).Info()
			if info&types.IsInteger == 0 {
				log.Fatalf("can't handle non-integer constant type %s", typ)
			}
			value := obj.(*types.Const).Val() // Guaranteed to succeed as this is CONST.
			if value.Kind() != constant.Int {
				log.Fatalf("can't happen: constant is not an integer %s", name)
			}
			i64, isInt := constant.Int64Val(value)
			u64, isUint := constant.Uint64Val(value)
			if !isInt && !isUint {
				log.Fatalf("internal error: value of %s is not an integer: %s", name, value.String())
			}
			if !isInt {
				u64 = uint64(i64)
			}
			v := Value{
				originalName: name.Name,
				value:        u64,
				signed:       info&types.IsUnsigned == 0,
				str:          value.String(),
			}
			if vspec.Doc != nil && vspec.Doc.Text() != "" {
				v.comment = vspec.Doc.Text()
			} else if c := vspec.Comment; c != nil && len(c.List) == 1 {
				v.comment = c.Text()
			}

			v.name = strings.TrimPrefix(v.originalName, f.trimPrefix)
			f.values = append(f.values, v)
		}
	}

	return false
}
