package main

import (
	"bytes"
	"fmt"
	"github.com/go-leo/gors/internal/pkg/gors"
	"go/ast"
	"io"
	"log"
	"strconv"
)

const (
	contextPackage = gors.GoImportPath("context")
	gorsPackage    = gors.GoImportPath("github.com/go-leo/gors")
	ginPackage     = gors.GoImportPath("github.com/gin-gonic/gin")
	httpPackage    = gors.GoImportPath("net/http")
	ioPackage      = gors.GoImportPath("io")
	bindingPackage = gors.GoImportPath("github.com/gin-gonic/gin/binding")
	renderPackage  = gors.GoImportPath("github.com/gin-gonic/gin/render")
	convxPackage   = gors.GoImportPath("github.com/go-leo/gox/convx")
)

type generate struct {
	buf              *bytes.Buffer
	headerBuf        *bytes.Buffer
	importsBuf       *bytes.Buffer
	functionBuf      *bytes.Buffer
	header           string
	pkgName          string
	imports          map[string]*gors.GoImport
	srvName          string
	usedPackageNames map[string]bool
	routerInfos      []*gors.RouterInfo
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

func (g *generate) checkAndGetResult1(rpcType *ast.FuncType, methodName *ast.Ident) *gors.Result {
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
		return &gors.Result{Bytes: true}
	case *ast.Ident:
		if r1.Name != "string" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &gors.Result{String: true}
	case *ast.StarExpr:
		switch x := r1.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &gors.Result{ObjectArgs: &gors.ObjectArgs{Name: name}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &gors.Result{ObjectArgs: &gors.ObjectArgs{Name: x.Sel.Name, GoImportPath: gors.GoImportPath(importPath)}}
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
		return &gors.Result{Reader: true}
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

func (g *generate) checkAndGetParam2(rpcType *ast.FuncType, methodName *ast.Ident) *gors.Param {
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
		return &gors.Param{Bytes: true}
	case *ast.Ident:
		if p2.Name != "string" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &gors.Param{String: true}
	case *ast.StarExpr:
		switch x := p2.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &gors.Param{ObjectArgs: &gors.ObjectArgs{Name: name}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &gors.Param{ObjectArgs: &gors.ObjectArgs{Name: x.Sel.Name, GoImportPath: gors.GoImportPath(importPath)}}
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
		return &gors.Param{Reader: true}
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
	functionName := serviceName + "Routes"
	g.P(g.functionBuf, "func ", functionName, "(srv ", serviceName, ", opts ...", gorsPackage.Ident("Option"), ") []", gorsPackage.Ident("Route"), " {")
	g.P(g.functionBuf, "options := ", gorsPackage.Ident("New"), "(opts...)")
	g.P(g.functionBuf, "_ = options")
	g.P(g.functionBuf, "return []", gorsPackage.Ident("Route"), "{")
	for _, routerInfo := range g.routerInfos {
		g.printRouterInfo(routerInfo)
	}
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "}")
}

func (g *generate) printRouterInfo(info *gors.RouterInfo) {
	g.P(g.functionBuf, gorsPackage.Ident("NewRoute"), "(")
	g.P(g.functionBuf, httpPackage.Ident(info.Method), ",")
	g.P(g.functionBuf, strconv.Quote(info.Path), ",")
	g.P(g.functionBuf, "func(c *", ginPackage.Ident("Context"), ") {")
	g.printHandler(info)
	g.P(g.functionBuf, "},")
	g.P(g.functionBuf, "),")
}

func (g *generate) printHandler(info *gors.RouterInfo) {
	fmName := fmt.Sprintf("/%s.%s/%s", g.pkgName, g.srvName, info.RpcMethodName)
	g.P(g.functionBuf, "var rpcMethodName = ", strconv.Quote(fmName))
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
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.RpcMethodName)
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
		log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or *struct{}", info.RpcMethodName)
	}
	g.P(g.functionBuf, "var err error")

	if info.Param2.Bytes {
		g.printBytesReq(info)
		g.P(g.functionBuf, "req = body")
	} else if info.Param2.String {
		g.printBytesReq(info)
		g.P(g.functionBuf, "req = string(body)")
	} else if info.Param2.Reader {
		g.P(g.functionBuf, "req = c.Request.Body")
	} else if info.Param2.ObjectArgs != nil {
		g.printObjectReq(info)
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.RpcMethodName)
	}

	g.P(g.functionBuf, "resp, err = srv.", info.RpcMethodName, "(ctx, req)")

	g.printResponseRender(info)

}

func (g *generate) printBytesReq(info *gors.RouterInfo) {
	g.P(g.functionBuf, "var body []byte")
	g.P(g.functionBuf, "body, err = ", ioPackage.Ident("ReadAll"), "(c.Request.Body)")
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printObjectReq(info *gors.RouterInfo) {
	objArgs := info.Param2.ObjectArgs
	g.P(g.functionBuf, "req = new(", objArgs.GoImportPath.Ident(objArgs.Name), ")")
	g.P(g.functionBuf, "if err = ", gorsPackage.Ident("RequestBind"), "(")
	g.P(g.functionBuf, "ctx, req, options.Tag,")
	for _, binding := range info.Bindings {
		g.P(g.functionBuf, gorsPackage.Ident(binding), ",")
	}
	g.P(g.functionBuf, "); err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printResponseRender(info *gors.RouterInfo) {
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")

	var renderMethodName string
	switch {
	case info.Result1.Bytes:
		switch {
		case info.BytesRender:
			renderMethodName = "BytesRender"
		default:
			log.Fatalf("error: func %s []byte result must be set BytesRender", info.RpcMethodName)
		}
	case info.Result1.String:
		switch {
		case info.StringRender:
			renderMethodName = "StringRender"
		case info.TextRender:
			renderMethodName = "TextRender"
		case info.HTMLRender:
			renderMethodName = "HTMLRender"
		case info.RedirectRender:
			renderMethodName = "RedirectRender"
		default:
			log.Fatalf("error: func %s string result must be set BytesRender or StringRender or TextRender or HTMLRender or RedirectRender", info.RpcMethodName)
		}
	case info.Result1.Reader:
		switch {
		case info.ReaderRender:
			renderMethodName = "ReaderRender"
		default:
			log.Fatalf("error: func %s io.Reader result must be set ReaderRender", info.RpcMethodName)
		}
	case info.Result1.ObjectArgs != nil:
		switch {
		case info.JSONRender:
			renderMethodName = "JSONRender"
		case info.IndentedJSONRender:
			renderMethodName = "IndentedJSONRender"
		case info.SecureJSONRender:
			renderMethodName = "SecureJSONRender"
		case info.JSONPJSONRender:
			renderMethodName = "JSONPJSONRender"
		case info.PureJSONRender:
			renderMethodName = "PureJSONRender"
		case info.AsciiJSONRender:
			renderMethodName = "AsciiJSONRender"
		case info.ProtoJSONRender:
			renderMethodName = "ProtoJSONRender"
		case info.XMLRender:
			renderMethodName = "XMLRender"
		case info.YAMLRender:
			renderMethodName = "YAMLRender"
		case info.ProtoBufRender:
			renderMethodName = "ProtoBufRender"
		case info.MsgPackRender:
			renderMethodName = "MsgPackRender"
		case info.TOMLRender:
			renderMethodName = "TOMLRender"
		case info.CustomRender:
			renderMethodName = "CustomRender"

		default:
			log.Fatalf("error: func %s *struct{} result must be set a Render", info.RpcMethodName)
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or *struct{}", info.RpcMethodName)
	}

	g.P(g.functionBuf, gorsPackage.Ident("ResponseRender"),
		"(ctx, ", gorsPackage.Ident("StatusCode"), "(ctx), resp,",
		strconv.Quote(info.RenderContentType), ",", gorsPackage.Ident(renderMethodName),
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
		case *gors.GoIdent:
			x.GoImport.Enable = true
			g.imports[x.GoImport.ImportPath] = x.GoImport
			_, _ = fmt.Fprint(w, x.Qualify())
		default:
			_, _ = fmt.Fprint(w, x)
		}
	}
	_, _ = fmt.Fprintln(w)
}
