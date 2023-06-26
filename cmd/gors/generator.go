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
)

type generate struct {
	buf              *bytes.Buffer
	headerBuf        *bytes.Buffer
	importsBuf       *bytes.Buffer
	functionBuf      *bytes.Buffer
	header           string
	pkg              string
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
	g.P(g.headerBuf, g.pkg)
}

func (g *generate) printFunction() {
	serviceName := g.srvName
	functionName := serviceName + "Routes"
	g.P(g.functionBuf, "func ", functionName, "(srv ", serviceName, ") []", gorsPackage.Ident("Route"), " {")
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
		g.printReqValidate()
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.RpcMethodName)
	}

	g.printRPCHandler(info)

	g.printResponse(info)

}

func (g *generate) printBytesReq(info *gors.RouterInfo) {
	g.P(g.functionBuf, "var body []byte")
	g.P(g.functionBuf, "body, err = ", ioPackage.Ident("ReadAll"), "(c.Request.Body)")
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("HandleBadRequest"), "(c, err)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printRequestBind(bindings []string) {
	g.P(g.functionBuf, "if err := ", gorsPackage.Ident("ShouldBind"), "(")
	g.P(g.functionBuf, "c, req, ")
	for _, binding := range bindings {
		g.P(g.functionBuf, gorsPackage.Ident(binding), ",")
	}
	g.P(g.functionBuf, "); err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("HandleBadRequest"), "(c, err)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printObjectReq(info *gors.RouterInfo) {
	objArgs := info.Param2.ObjectArgs
	g.P(g.functionBuf, "req = new(", objArgs.GoImportPath.Ident(objArgs.Name), ")")
	var bindings []string
	if info.UriBinding {
		bindings = append(bindings, "UriBinding")
	}
	if info.QueryBinding {
		bindings = append(bindings, "QueryBinding")
	}
	if info.HeaderBinding {
		bindings = append(bindings, "HeaderBinding")
	}
	if info.FormBinding {
		bindings = append(bindings, "FormBinding")
	}
	if info.FormPostBinding {
		bindings = append(bindings, "FormPostBinding")
	}
	if info.FormMultipartBinding {
		bindings = append(bindings, "FormMultipartBinding")
	}
	if info.JSONBinding {
		bindings = append(bindings, "JSONBinding")
	}
	if info.XMLBinding {
		bindings = append(bindings, "XMLBinding")
	}
	if info.ProtoBufBinding {
		bindings = append(bindings, "ProtoBufBinding")
	}
	if info.MsgPackBinding {
		bindings = append(bindings, "MsgPackBinding")
	}
	if info.YAMLBinding {
		bindings = append(bindings, "YAMLBinding")
	}
	if info.TOMLBinding {
		bindings = append(bindings, "TOMLBinding")
	}
	if info.CustomBinding {
		bindings = append(bindings, "CustomBinding")
	}
	g.printRequestBind(bindings)
}

func (g *generate) printReqValidate() {
	g.P(g.functionBuf, "if err = ", gorsPackage.Ident("Validate"), "(req); err != nil {")
	g.P(g.functionBuf, gorsPackage.Ident("HandleBadRequest"), "(c, err)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printRPCHandler(info *gors.RouterInfo) {
	g.P(g.functionBuf, "ctx := ", gorsPackage.Ident("NewContext"), "(c)")
	g.P(g.functionBuf, "resp, err = srv.", info.RpcMethodName, "(ctx, req)")
}

func (g *generate) printResponse(info *gors.RouterInfo) {
	g.P(g.functionBuf, "switch e := err.(type) {")
	g.P(g.functionBuf, "case nil:")

	switch {
	case info.Result1.Bytes:
		switch {
		case info.BytesRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("Data"), "{ContentType: ", strconv.Quote(info.RenderContentType), ", Data: resp})")
		default:
			log.Fatalf("error: func %s []byte result must be set BytesRender", info.RpcMethodName)
		}
	case info.Result1.String:
		switch {
		case info.StringRender, info.TextRender, info.HTMLRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("Data"), "{ContentType: ", strconv.Quote(info.RenderContentType), ", Data: []byte(resp)})")
		case info.RedirectRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.Redirect(statusCode, resp)")
		default:
			log.Fatalf("error: func %s string result must be set BytesRender or StringRender or TextRender or HTMLRender or RedirectRender", info.RpcMethodName)
		}
	case info.Result1.Reader:
		switch {
		case info.ReaderRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("Reader"), "{ContentType: ", strconv.Quote(info.RenderContentType), ", ContentLength: -1, Reader: resp})")
		default:
			log.Fatalf("error: func %s io.Reader result must be set ReaderRender", info.RpcMethodName)
		}
	case info.Result1.ObjectArgs != nil:
		switch {
		case info.JSONRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.JSON(statusCode, resp)")
		case info.IndentedJSONRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.IndentedJSON(statusCode, resp)")
		case info.SecureJSONRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.SecureJSON(statusCode, resp)")
		case info.JsonpJSONRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.JSONP(statusCode, resp)")
		case info.PureJSONRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.PureJSON(statusCode, resp)")
		case info.AsciiJSONRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.AsciiJSON(statusCode, resp)")
		case info.XMLRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.XML(statusCode, resp)")
		case info.YAMLRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.YAML(statusCode, resp)")
		case info.ProtobufRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.ProtoBuf(statusCode, resp)")
		case info.MsgPackRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("MsgPack"), "{Data: resp})")
		case info.TOMLRender:
			g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("HttpStatusCode"), "(c, resp)")
			g.P(g.functionBuf, "c.TOML(statusCode, resp)")
		case info.CustomRender:
			g.P(g.functionBuf, "var render ", gorsPackage.Ident("Render"), " = resp")
			g.P(g.functionBuf, "render.Render(c)")
		default:
			log.Fatalf("error: func %s *struct{} result must be set a Render", info.RpcMethodName)
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or *struct{}", info.RpcMethodName)
	}

	g.P(g.functionBuf, "return")
	//g.P(g.functionBuf, "case *", gorsPackage.Ident("Error"), ":")
	//g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "case *", gorsPackage.Ident("HttpError"), ":")
	g.P(g.functionBuf, "c.String(e.StatusCode(), e.Error())")
	g.P(g.functionBuf, "_ = c.Error(e).SetType(", ginPackage.Ident("ErrorTypePublic"), ")")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "default:")
	g.P(g.functionBuf, "c.String(", httpPackage.Ident("StatusInternalServerError"), ", err.Error())")
	g.P(g.functionBuf, "_ = c.Error(e).SetType(", ginPackage.Ident("ErrorTypePrivate"), ")")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")

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
