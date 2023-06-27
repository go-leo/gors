package main

import (
	//"strings"

	"bufio"
	"fmt"
	"github.com/go-leo/gors/internal/pkg/gors"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
	"strings"
)

const (
	metadataPackage = protogen.GoImportPath("google.golang.org/grpc/metadata")
	grpcPackage     = protogen.GoImportPath("google.golang.org/grpc")
	ginPackage      = protogen.GoImportPath("github.com/gin-gonic/gin")
	httpPackage     = protogen.GoImportPath("net/http")
	gorsPackage     = protogen.GoImportPath("github.com/go-leo/gors")
	bindingPackage  = protogen.GoImportPath("github.com/gin-gonic/gin/binding")
)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_grpc.gors.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-gors. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	for _, service := range file.Services {
		genClientFunction(gen, file, g, service)
		genServerFunction(gen, file, g, service)
	}
}

func genClientFunction(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	clientName := service.GoName + "Client"
	funcName := clientName + "Routes"

	basePath := extractBasePath(service)
	g.P("func ", funcName, "(cli ", clientName, ") []", gorsPackage.Ident("Route"), " {")
	g.P("return []", gorsPackage.Ident("Route"), "{")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			// Unary RPC method
			fmName := fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())

			router := newRouter(method, basePath)
			g.P(gorsPackage.Ident("NewRoute"), "(")
			g.P(httpPackage.Ident(router.Method), ",")
			g.P(strconv.Quote(router.Path), ",")
			g.P("func(c *", ginPackage.Ident("Context"), ") {")
			g.P("var req *", method.Input.GoIdent)
			g.P("var resp *", method.Output.GoIdent)
			g.P("var err error")
			g.P("req = new(", method.Input.GoIdent, ")")

			printRequestBinding(gen, g, router, fmName)

			g.P("ctx := ", gorsPackage.Ident("NewContext"), "(c)")
			g.P("resp, err = cli.", method.GoName, "(ctx, req)")

			printResponseRender(gen, g, router, fmName)

			g.P("},")
			g.P("),")
			//
			//	var headerMD, trailerMD metadata.MD
			//	resp, err := cli.SayHello(ctx, req, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
			//	grpcproxy.Render(c, headerMD, trailerMD, resp, err)
		} else {
			// Streaming RPC method
			continue
		}
	}
	g.P("}")
	g.P("}")
}

func genServerFunction(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	serverName := service.GoName + "Server"
	funcName := serverName + "Routes"

	basePath := extractBasePath(service)
	g.P("func ", funcName, "(srv ", serverName, ") []", gorsPackage.Ident("Route"), " {")
	g.P("return []", gorsPackage.Ident("Route"), "{")
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			// Unary RPC method
			fmName := fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())

			router := newRouter(method, basePath)
			g.P(gorsPackage.Ident("NewRoute"), "(")
			g.P(httpPackage.Ident(router.Method), ",")
			g.P(strconv.Quote(router.Path), ",")
			g.P("func(c *", ginPackage.Ident("Context"), ") {")
			g.P("var req *", method.Input.GoIdent)
			g.P("var resp *", method.Output.GoIdent)
			g.P("var err error")
			g.P("req = new(", method.Input.GoIdent, ")")

			printRequestBinding(gen, g, router, fmName)

			g.P("ctx := ", gorsPackage.Ident("NewContext"), "(c)")
			g.P("resp, err = srv.", method.GoName, "(ctx, req)")

			printResponseRender(gen, g, router, fmName)

			g.P("},")
			g.P("),")
			//
			//	var headerMD, trailerMD metadata.MD
			//	resp, err := cli.SayHello(ctx, req, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
			//	grpcproxy.Render(c, headerMD, trailerMD, resp, err)
		} else {
			// Streaming RPC method
			continue
		}
	}
	g.P("}")
	g.P("}")
}

func printRequestBinding(gen *protogen.Plugin, g *protogen.GeneratedFile, router *gors.RouterInfo, fmName string) {
	var bindings []string
	if router.UriBinding {
		bindings = append(bindings, "UriBindingWith")
	}
	if router.QueryBinding {
		bindings = append(bindings, "QueryBindingWith")
	}
	if router.HeaderBinding {
		bindings = append(bindings, "HeaderBindingWith")
	}
	if router.FormBinding {
		bindings = append(bindings, "FormBindingWith")
	}
	if router.FormPostBinding {
		bindings = append(bindings, "FormPostBindingWith")
	}
	if router.JSONBinding {
		bindings = append(bindings, "JSONBindingWith")
	}
	if router.ProtoBufBinding {
		bindings = append(bindings, "ProtoBufBindingWith")
	}
	if router.CustomBinding {
		bindings = append(bindings, "CustomBindingWith")
	}
	if router.MsgPackBinding {
		bindings = append(bindings, "MsgPackBindingWith")
	}
	if router.XMLBinding {
		gen.Error(fmt.Errorf("%s, @XMLBinding is not supported", fmName))
		return
	}
	if router.YAMLBinding {
		gen.Error(fmt.Errorf("%s, @YAMLBinding is not supported", fmName))
		return
	}
	if router.TOMLBinding {
		gen.Error(fmt.Errorf("%s, @TOMLBinding is not supported", fmName))
		return
	}
	g.P("if err = ", gorsPackage.Ident("ShouldBindWith"), "(")
	g.P("c, req, ", strconv.Quote("json"), ",")
	for _, binding := range bindings {
		g.P(gorsPackage.Ident(binding), ",")
	}
	g.P("); err != nil {")
	g.P(gorsPackage.Ident("HandleBadRequest"), "(c, err)")
	g.P("return")
	g.P("}")
}

func printResponseRender(gen *protogen.Plugin, g *protogen.GeneratedFile, router *gors.RouterInfo, fmName string) {
	switch {
	case router.JSONRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("JSONRender"), ")")
	case router.IndentedJSONRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("IndentedJSONRender"), ")")
	case router.SecureJSONRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("SecureJSONRender"), ")")
	case router.JSONPJSONRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("JSONPJSONRender"), ")")
	case router.PureJSONRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("PureJSONRender"), ")")
	case router.AsciiJSONRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("AsciiJSONRender"), ")")
	case router.ProtoBufRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("ProtoBufRender"), ")")
	case router.MsgPackRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("MsgPackRender"), ")")
	case router.CustomRender:
		g.P(gorsPackage.Ident("MustRender"), "(c, resp, err, ", strconv.Quote(router.RenderContentType), ", ", gorsPackage.Ident("CustomRender"), ")")
	case router.BytesRender:
		gen.Error(fmt.Errorf("%s, @BytesRender is not supported", fmName))
		return
	case router.StringRender:
		gen.Error(fmt.Errorf("%s, @StringRender is not supported", fmName))
		return
	case router.TextRender:
		gen.Error(fmt.Errorf("%s, @TextRender is not supported", fmName))
		return
	case router.HTMLRender:
		gen.Error(fmt.Errorf("%s, @HTMLRender is not supported", fmName))
		return
	case router.RedirectRender:
		gen.Error(fmt.Errorf("%s, @RedirectRender is not supported", fmName))
		return
	case router.ReaderRender:
		gen.Error(fmt.Errorf("%s, @ReaderRender is not supported", fmName))
		return
	case router.XMLRender:
		gen.Error(fmt.Errorf("%s, @XMLRender is not supported", fmName))
		return
	case router.YAMLRender:
		gen.Error(fmt.Errorf("%s, @YAMLRender is not supported", fmName))
		return
	case router.TOMLRender:
		gen.Error(fmt.Errorf("%s, @TOMLRender is not supported", fmName))
		return
	default:
		gen.Error(fmt.Errorf("%s, render not defined", fmName))
		return
	}
}

func extractBasePath(service *protogen.Service) string {
	return gors.ExtractBasePath(splitComment(service.Comments.Leading.String()))
}

func newRouter(method *protogen.Method, basePath string) *gors.RouterInfo {
	return gors.NewRouter(method.GoName, basePath, splitComment(method.Comments.Leading.String()))
}

func splitComment(leadingComment string) []string {
	var comments []string
	scanner := bufio.NewScanner(strings.NewReader(leadingComment))
	for scanner.Scan() {
		line := scanner.Text()
		comments = append(comments, line)
	}
	return comments
}
