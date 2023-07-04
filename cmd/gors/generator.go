package main

import (
	"bytes"
	"fmt"
	"github.com/go-leo/gors/internal/pkg/annotation"
	"github.com/go-leo/gox/stringx"
	"go/ast"
	"golang.org/x/exp/slices"
	"io"
	"log"
	"path"
	"strconv"
	"strings"
)

const (
	contextPackage = annotation.GoImportPath("context")
	gorsPackage    = annotation.GoImportPath("github.com/go-leo/gors")
	ginPackage     = annotation.GoImportPath("github.com/gin-gonic/gin")
	httpPackage    = annotation.GoImportPath("net/http")
	ioPackage      = annotation.GoImportPath("io")
	bindingPackage = annotation.GoImportPath("github.com/gin-gonic/gin/binding")
	renderPackage  = annotation.GoImportPath("github.com/gin-gonic/gin/render")
	convxPackage   = annotation.GoImportPath("github.com/go-leo/gox/convx")
)

type generate struct {
	buf              *bytes.Buffer
	headerBuf        *bytes.Buffer
	importsBuf       *bytes.Buffer
	functionBuf      *bytes.Buffer
	header           string
	pkgName          string
	imports          map[string]*annotation.GoImport
	srvName          string
	usedPackageNames map[string]bool
	routerInfos      []*annotation.RouterInfo
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

func (g *generate) checkAndGetResult1(rpcType *ast.FuncType, methodName *ast.Ident) *annotation.Result {
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
		return &annotation.Result{Bytes: true}
	case *ast.Ident:
		if r1.Name != "string" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &annotation.Result{String: true}
	case *ast.StarExpr:
		switch x := r1.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &annotation.Result{ObjectArgs: &annotation.ObjectArgs{Name: name}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &annotation.Result{ObjectArgs: &annotation.ObjectArgs{Name: x.Sel.Name, GoImportPath: annotation.GoImportPath(importPath)}}
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
		return &annotation.Result{Reader: true}
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

func (g *generate) checkAndGetParam2(rpcType *ast.FuncType, methodName *ast.Ident) *annotation.Param {
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
		return &annotation.Param{Bytes: true}
	case *ast.Ident:
		if p2.Name != "string" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &annotation.Param{String: true}
	case *ast.StarExpr:
		switch x := p2.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &annotation.Param{ObjectArgs: &annotation.ObjectArgs{Name: name}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &annotation.Param{ObjectArgs: &annotation.ObjectArgs{Name: x.Sel.Name, GoImportPath: annotation.GoImportPath(importPath)}}
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
		return &annotation.Param{Reader: true}
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
	serviceName := g.srvName

	for _, routerInfo := range g.routerInfos {
		g.printRouterMethod(serviceName, routerInfo)
	}

	functionName := serviceName + "Routes"
	g.P(g.functionBuf, "func ", functionName, "(srv ", serviceName, ", opts ...", gorsPackage.Ident("Option"), ") []", gorsPackage.Ident("Route"), " {")
	g.P(g.functionBuf, "options := ", gorsPackage.Ident("New"), "(opts...)")
	g.P(g.functionBuf, "return []", gorsPackage.Ident("Route"), "{")
	for _, routerInfo := range g.routerInfos {
		g.printRouterInfo(serviceName, routerInfo)
	}
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "}")
}

func (g *generate) printRouterMethod(serviceName string, info *annotation.RouterInfo) {
	handlerName := fmt.Sprintf("_%s_%s_Handler", g.srvName, info.MethodName)
	info.HandlerName = handlerName
	g.P(g.functionBuf, "func ", handlerName, "(srv ", serviceName, ", options *", gorsPackage.Ident("Options"), ")", "func(c *", ginPackage.Ident("Context"), ") {")
	g.P(g.functionBuf, "return func(c *", ginPackage.Ident("Context"), ") {")
	g.printHandler(info)
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "")
}

func (g *generate) printRouterInfo(serviceName string, info *annotation.RouterInfo) {
	p := path.Join(info.BasePath, info.Path)
	g.P(g.functionBuf, gorsPackage.Ident("NewRoute"), "(", httpPackage.Ident(info.HttpMethod.HttpMethod()), ",", strconv.Quote(p), ",", info.HandlerName, "(srv, options),", "),")
}

func (g *generate) printHandler(info *annotation.RouterInfo) {
	g.P(g.functionBuf, "var rpcMethodName = ", strconv.Quote(info.FullMethodName))
	g.P(g.functionBuf, "var ctx = ", gorsPackage.Ident("NewContext"), "(c, rpcMethodName)")

	if info.Param2.Bytes {
		g.P(g.functionBuf, "var req []byte")
	} else if info.Param2.String {
		g.P(g.functionBuf, "var req string")
	} else if info.Param2.Reader {
		g.P(g.functionBuf, "var req ", ioPackage.Ident("Reader"))
	} else if objectArgs := info.Param2.ObjectArgs; objectArgs != nil {
		g.P(g.functionBuf, "var req *", objectArgs.GoImportPath.Ident(objectArgs.Name))
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}
	if info.Result1.Bytes {
		g.P(g.functionBuf, "var resp []byte")
	} else if info.Result1.String {
		g.P(g.functionBuf, "var resp string")
	} else if info.Result1.Reader {
		g.P(g.functionBuf, "var resp ", ioPackage.Ident("Reader"))
	} else if objectArgs := info.Result1.ObjectArgs; objectArgs != nil {
		g.P(g.functionBuf, "var resp *", objectArgs.GoImportPath.Ident(objectArgs.Name))
	} else {
		log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}
	g.P(g.functionBuf, "var err error")

	if info.Param2.Reader {
		g.printPtrReq(info, annotation.ReaderBinding)
	} else if info.Param2.Bytes {
		g.printPtrReq(info, annotation.BytesBinding)
	} else if info.Param2.String {
		g.printPtrReq(info, annotation.StringBinding)
	} else if info.Param2.ObjectArgs != nil {
		g.printObjectReq(info)
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}

	g.P(g.functionBuf, "resp, err = srv.", info.MethodName, "(ctx, req)")

	g.printResponseRender(info)

}

func (g *generate) printPtrReq(info *annotation.RouterInfo, binding string) {
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

func (g *generate) printObjectReq(info *annotation.RouterInfo) {
	objArgs := info.Param2.ObjectArgs
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

func (g *generate) printResponseRender(info *annotation.RouterInfo) {
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")

	switch {
	case info.Result1.Bytes:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.BytesRender
		}
		if info.Render != annotation.BytesRender {
			log.Fatalf("error: func %s []byte result must be set %s", info.FullMethodName, annotation.BytesRender)
			return
		}
	case info.Result1.String:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.StringRender
		}
		renders := []string{annotation.StringRender, annotation.TextRender, annotation.HTMLRender, annotation.RedirectRender}
		if !slices.Contains(renders, info.Render) {
			log.Fatalf("error: func %s string result must be set %v", info.FullMethodName, renders)
		}
	case info.Result1.Reader:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.ReaderRender
		}
		if info.Render != annotation.ReaderRender {
			log.Fatalf("error: func %s io.Reader result must be set %s", info.FullMethodName, annotation.ReaderRender)
			return
		}
	case info.Result1.ObjectArgs != nil:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.JSONRender
			info.RenderContentType = annotation.JSONContentType
		}
		renders := []string{
			annotation.JSONRender, annotation.IndentedJSONRender, annotation.SecureJSONRender, annotation.JSONPJSONRender,
			annotation.PureJSONRender, annotation.AsciiJSONRender, annotation.ProtoJSONRender, annotation.XMLRender,
			annotation.YAMLRender, annotation.ProtoBufRender, annotation.MsgPackRender, annotation.TOMLRender,
			annotation.CustomRender,
		}
		if !slices.Contains(renders, info.Render) {
			log.Fatalf("error: func %s *struct result must be set %v", info.FullMethodName, renders)
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be io.Reader or []byte or string or *struct{}", info.FullMethodName)
	}

	g.P(g.functionBuf, gorsPackage.Ident("ResponseRender"),
		"(ctx, ", gorsPackage.Ident("StatusCode"), "(ctx), resp,",
		strconv.Quote(info.RenderContentType), ",", gorsPackage.Ident(strings.TrimPrefix(info.Render, "@")),
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
		case *annotation.GoIdent:
			x.GoImport.Enable = true
			g.imports[x.GoImport.ImportPath] = x.GoImport
			_, _ = fmt.Fprint(w, x.Qualify())
		default:
			_, _ = fmt.Fprint(w, x)
		}
	}
	_, _ = fmt.Fprintln(w)
}
