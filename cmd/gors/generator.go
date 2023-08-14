package main

import (
	"bytes"
	"fmt"
	"github.com/go-leo/gors/internal/pkg/parser"
	"github.com/go-leo/gox/stringx"
	"go/ast"
	"go/token"
	"golang.org/x/exp/slices"
	"io"
	"log"
	"path"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	contextPackage = GoImportPath("context")
	gorsPackage    = GoImportPath("github.com/go-leo/gors")
	ginPackage     = GoImportPath("github.com/gin-gonic/gin")
	httpPackage    = GoImportPath("net/http")
	ioPackage      = GoImportPath("io")
	bindingPackage = GoImportPath("github.com/gin-gonic/gin/binding")
	renderPackage  = GoImportPath("github.com/gin-gonic/gin/render")
	convxPackage   = GoImportPath("github.com/go-leo/gox/convx")
)

type GoImportPath string

func (p GoImportPath) Ident(s string) *GoIdent {
	importPath := string(p)
	return &GoIdent{
		GoName: s,
		GoImport: &GoImport{
			PackageName: CleanPackageName(path.Base(importPath)),
			ImportPath:  importPath,
		},
	}
}

type GoIdent struct {
	GoImport *GoImport
	GoName   string
}

func (x *GoIdent) Qualify() string {
	if x.GoImport.ImportPath == "" {
		return x.GoName
	}
	return x.GoImport.PackageName + "." + x.GoName
}

type GoImport struct {
	PackageName string
	ImportPath  string
	Enable      bool
}

type ObjectArgs struct {
	Name         string
	GoImportPath GoImportPath
	StarExpr     *ast.StarExpr
}

type Param struct {
	Bytes      bool
	String     bool
	ObjectArgs *ObjectArgs
	Reader     bool
}

type Result struct {
	Bytes      bool
	String     bool
	ObjectArgs *ObjectArgs
	Reader     bool
}

type generate struct {
	buf              *bytes.Buffer
	headerBuf        *bytes.Buffer
	importsBuf       *bytes.Buffer
	functionBuf      *bytes.Buffer
	header           string
	pkgName          string
	imports          map[string]*GoImport
	usedPackageNames map[string]bool
	serviceInfo      *parser.ServiceInfo
	Param2s          map[*parser.RouterInfo]*Param
	Result1s         map[*parser.RouterInfo]*Result
}

func (g *generate) checkResult2MustBeError(rpcType *ast.FuncType, methodName *ast.Ident) {
	result2 := rpcType.Results.List[1]
	result2Iden, ok := result2.Type.(*ast.Ident)
	if !ok {
		log.Fatalf("error: func %s 2th result is not error", methodName)
	}
	if result2Iden.Name != "error" {
		log.Fatalf("error: func %s 2th result is not error", methodName)
	}
}

func (g *generate) checkAndGetResult1(rpcType *ast.FuncType, methodName *ast.Ident) *Result {
	result1 := rpcType.Results.List[0]
	switch r1 := result1.Type.(type) {
	case *ast.ArrayType:
		ident, ok := r1.Elt.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ident.Name != "byte" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &Result{Bytes: true}
	case *ast.Ident:
		if r1.Name != "string" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &Result{String: true}
	case *ast.StarExpr:
		switch x := r1.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &Result{
				ObjectArgs: &ObjectArgs{
					Name:     name,
					StarExpr: r1,
				}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &Result{
						ObjectArgs: &ObjectArgs{
							Name:         x.Sel.Name,
							GoImportPath: GoImportPath(importPath),
							StarExpr:     r1,
						}}
				}
			}
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		default:
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		}
	case *ast.SelectorExpr:
		if r1.Sel == nil {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if r1.Sel.Name != "Reader" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ident, ok := r1.X.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ioImport, ok := g.imports["io"]
		if !ok {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ioImport.PackageName != ident.Name {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &Result{Reader: true}
	default:

	}
	return nil
}

func (g *generate) checkResults(rpcType *ast.FuncType, methodName *ast.Ident) {
	if rpcType.Results == nil {
		log.Fatalf("error: func %s results is empty", methodName)
	}
	if len(rpcType.Results.List) != 2 {
		log.Fatalf("error: func %s results count is not equal 2", methodName)
	}
}

func (g *generate) checkAndGetParam2(rpcType *ast.FuncType, methodName *ast.Ident) *Param {
	param2 := rpcType.Params.List[1]
	switch p2 := param2.Type.(type) {
	case *ast.ArrayType:
		ident, ok := p2.Elt.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ident.Name != "byte" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &Param{Bytes: true}
	case *ast.Ident:
		if p2.Name != "string" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &Param{String: true}
	case *ast.StarExpr:
		switch x := p2.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &Param{
				ObjectArgs: &ObjectArgs{
					Name:     name,
					StarExpr: p2,
				},
			}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &Param{
						ObjectArgs: &ObjectArgs{
							Name:         x.Sel.Name,
							GoImportPath: GoImportPath(importPath),
							StarExpr:     p2,
						},
					}
				}
			}
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		default:
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		}

	case *ast.SelectorExpr:
		if p2.Sel == nil {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if p2.Sel.Name != "Reader" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ident, ok := p2.X.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ioImport, ok := g.imports["io"]
		if !ok {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ioImport.PackageName != ident.Name {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &Param{Reader: true}
	default:
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		return nil
	}
}

func (g *generate) checkParam1MustBeContext(rpcType *ast.FuncType, methodName *ast.Ident) {
	param1 := rpcType.Params.List[0]
	param0SelectorExpr, ok := param1.Type.(*ast.SelectorExpr)
	if !ok {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
	if param0SelectorExpr.Sel.Name != "Context" {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
	param0SelectorExprX, ok := param0SelectorExpr.X.(*ast.Ident)
	if !ok {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
	if param0SelectorExprX.Name != "context" {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
}

func (g *generate) checkParams(rpcType *ast.FuncType, methodName *ast.Ident) {
	if rpcType.Params == nil {
		log.Fatalf("error: func %s params is empty", methodName)
	}
	if len(rpcType.Params.List) != 2 {
		log.Fatalf("error: func %s params count is not equal 2", methodName)
	}
}

func (g *generate) content() []byte {
	// header
	g.printHeader()

	// function
	g.printFunction()

	// imports
	g.printImports()

	// all
	g.combine()
	return g.buf.Bytes()
}

func (g *generate) printHeader() {
	g.P(g.headerBuf, g.header)
	g.P(g.headerBuf)
	g.P(g.headerBuf, fmt.Sprintf("package %s", g.pkgName))
}

func (g *generate) printFunction() {
	g.printRoutesMethod()
	g.printHandlerMethods()
}

func (g *generate) printRoutesMethod() {
	functionName := g.serviceInfo.Name + "Routes"
	g.P(g.functionBuf, "func ", functionName, "(srv ", g.serviceInfo.Name, ", opts ...", gorsPackage.Ident("Option"), ") []", gorsPackage.Ident("Route"), " {")
	g.P(g.functionBuf, "options := ", gorsPackage.Ident("NewOptions"), "(opts...)")
	g.P(g.functionBuf, "return []", gorsPackage.Ident("Route"), "{")
	for _, routerInfo := range g.serviceInfo.Routers {
		g.printRouterInfo(routerInfo)
	}
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "")
}

func (g *generate) printHandlerMethods() {
	for _, routerInfo := range g.serviceInfo.Routers {
		g.printHandlerMethod(routerInfo)
	}
}

func (g *generate) printHandlerMethod(info *parser.RouterInfo) {
	g.P(g.functionBuf, "func ", info.HandlerName, "(srv ", g.serviceInfo.Name, ", options *", gorsPackage.Ident("Options"), ")", "func(c *", ginPackage.Ident("Context"), ") {")
	g.P(g.functionBuf, "return func(c *", ginPackage.Ident("Context"), ") {")
	g.printHandler(info)
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "")
}

func (g *generate) printRouterInfo(info *parser.RouterInfo) {
	p := path.Join(g.serviceInfo.BasePath, info.Path)
	g.P(g.functionBuf, gorsPackage.Ident("NewRoute"), "(", httpPackage.Ident(info.HttpMethod.HttpMethod()), ",", strconv.Quote(p), ",", info.HandlerName, "(srv, options),", "),")
}

func (g *generate) printHandler(info *parser.RouterInfo) {
	g.P(g.functionBuf, "var rpcMethodName = ", strconv.Quote(info.FullMethodName))
	g.P(g.functionBuf, "var ctx = ", gorsPackage.Ident("NewContext"), "(c, rpcMethodName)")

	Param2 := g.Param2s[info]
	Result1 := g.Result1s[info]

	if Param2.Bytes {
		g.P(g.functionBuf, "var req []byte")
	} else if Param2.String {
		g.P(g.functionBuf, "var req string")
	} else if Param2.Reader {
		g.P(g.functionBuf, "var req ", ioPackage.Ident("Reader"))
	} else if objectArgs := Param2.ObjectArgs; objectArgs != nil {
		g.P(g.functionBuf, "var req *", objectArgs.GoImportPath.Ident(objectArgs.Name))
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}
	if Result1.Bytes {
		g.P(g.functionBuf, "var resp []byte")
	} else if Result1.String {
		g.P(g.functionBuf, "var resp string")
	} else if Result1.Reader {
		g.P(g.functionBuf, "var resp ", ioPackage.Ident("Reader"))
	} else if objectArgs := Result1.ObjectArgs; objectArgs != nil {
		g.P(g.functionBuf, "var resp *", objectArgs.GoImportPath.Ident(objectArgs.Name))
	} else {
		log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}
	g.P(g.functionBuf, "var err error")

	if Param2.Reader {
		g.printPtrReq(info, parser.ReaderBinding)
	} else if Param2.Bytes {
		g.printPtrReq(info, parser.BytesBinding)
	} else if Param2.String {
		g.printPtrReq(info, parser.StringBinding)
	} else if Param2.ObjectArgs != nil {
		g.printObjectReq(info)
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}

	g.P(g.functionBuf, "resp, err = srv.", info.MethodName, "(ctx, req)")

	g.printResponseRender(info)

}

func (g *generate) printPtrReq(info *parser.RouterInfo, binding string) {
	if len(info.Bindings) != 1 {
		log.Fatalf("error: binding must be %s", binding)
	}
	g.P(g.functionBuf, "if err = ", gorsPackage.Ident("RequestBind"), "(")
	g.P(g.functionBuf, "ctx, &req, options.Tag,")
	g.P(g.functionBuf, gorsPackage.Ident(strings.Trim(info.Bindings[0], "@")), ",")
	g.P(g.functionBuf, "); err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printObjectReq(info *parser.RouterInfo) {
	Param2 := g.Param2s[info]
	objArgs := Param2.ObjectArgs
	g.P(g.functionBuf, "req = new(", objArgs.GoImportPath.Ident(objArgs.Name), ")")
	g.P(g.functionBuf, "if err = ", gorsPackage.Ident("RequestBind"), "(")
	g.P(g.functionBuf, "ctx, req, options.Tag,")
	for _, binding := range info.Bindings {
		g.P(g.functionBuf, gorsPackage.Ident(strings.Trim(binding, "@")), ",")
	}
	g.P(g.functionBuf, "); err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printResponseRender(info *parser.RouterInfo) {
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")

	Result1 := g.Result1s[info]
	switch {
	case Result1.Bytes:
		if stringx.IsBlank(info.Render) {
			info.Render = parser.BytesRender
		}
		if info.Render != parser.BytesRender {
			log.Fatalf("error: func %s []byte result must be set %s", info.FullMethodName, parser.BytesRender)
			return
		}
	case Result1.String:
		if stringx.IsBlank(info.Render) {
			info.Render = parser.StringRender
		}
		renders := []string{parser.StringRender, parser.TextRender, parser.HTMLRender, parser.RedirectRender}
		if !slices.Contains(renders, info.Render) {
			log.Fatalf("error: func %s string result must be set %v", info.FullMethodName, renders)
		}
	case Result1.Reader:
		if stringx.IsBlank(info.Render) {
			info.Render = parser.ReaderRender
		}
		if info.Render != parser.ReaderRender {
			log.Fatalf("error: func %s io.Reader result must be set %s", info.FullMethodName, parser.ReaderRender)
			return
		}
	case Result1.ObjectArgs != nil:
		if stringx.IsBlank(info.Render) {
			info.Render = parser.JSONRender
			info.RenderContentType = parser.JSONContentType
		}
		renders := []string{
			parser.JSONRender, parser.IndentedJSONRender, parser.SecureJSONRender, parser.JSONPJSONRender,
			parser.PureJSONRender, parser.AsciiJSONRender, parser.ProtoJSONRender, parser.XMLRender,
			parser.YAMLRender, parser.ProtoBufRender, parser.MsgPackRender, parser.TOMLRender,
			parser.CustomRender,
		}
		if !slices.Contains(renders, info.Render) {
			log.Fatalf("error: func %s *struct result must be set %v", info.FullMethodName, renders)
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be io.Reader or []byte or string or *struct{}", info.FullMethodName)
	}

	renderName := strings.TrimPrefix(info.Render, "@")
	renderArg := ""
	if info.Render == parser.ProtoJSONRender {
		renderArg = "(options.ProtoJSONMarshalOptions)"
	}

	g.P(g.functionBuf, gorsPackage.Ident("ResponseRender"),
		"(ctx, ", gorsPackage.Ident("StatusCode"), "(ctx), resp,",
		strconv.Quote(info.RenderContentType), ",",
		gorsPackage.Ident(renderName), renderArg,
		", options.ResponseWrapper)")
}

func (g *generate) printImports() {
	g.P(g.importsBuf, "import (")
	for _, imp := range g.imports {
		if !imp.Enable {
			continue
		}
		if imp.ImportPath == "" {
			continue
		}
		g.P(g.importsBuf, imp.PackageName, " ", strconv.Quote(imp.ImportPath))
	}
	g.P(g.importsBuf, ")")

}

func (g *generate) combine() {
	_, _ = io.Copy(g.buf, g.headerBuf)
	_, _ = io.Copy(g.buf, g.importsBuf)
	_, _ = io.Copy(g.buf, g.functionBuf)
}

func (g *generate) P(w io.Writer, v ...any) {
	for _, x := range v {
		switch x := x.(type) {
		case *GoIdent:
			x.GoImport.Enable = true
			g.imports[x.GoImport.ImportPath] = x.GoImport
			_, _ = fmt.Fprint(w, x.Qualify())
		default:
			_, _ = fmt.Fprint(w, x)
		}
	}
	_, _ = fmt.Fprintln(w)
}

func (g *generate) doc() []byte {
	//

	return nil
}

func CleanPackageName(name string) string {
	name = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, name)

	// Prepend '_' in the event of a Go keyword conflict or if
	// the identifier is invalid (does not start in the Unicode L category).
	r, _ := utf8.DecodeRuneInString(name)
	if token.Lookup(name).IsKeyword() || !unicode.IsLetter(r) {
		return "_" + name
	}
	return name
}

func handlerName(info *parser.RouterInfo, serviceInfo *parser.ServiceInfo) string {
	return fmt.Sprintf("_%s_%s_Handler", serviceInfo.Name, info.MethodName)
}
