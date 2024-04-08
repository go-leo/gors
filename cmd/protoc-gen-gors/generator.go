package main

import (
	//"strings"

	"bufio"
	"fmt"
	"github.com/go-leo/gors/internal/pkg/parser"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/gox/stringx"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/compiler/protogen"
	"path"
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
	contextPackage  = protogen.GoImportPath("context")
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
		serviceInfo, err := getServiceInfo(gen, file, g, service)
		if err != nil {
			gen.Error(fmt.Errorf("error: %w", err))
			return
		}

		if err := genClientRoutesFunction(gen, file, g, service, serviceInfo); err != nil {
			gen.Error(fmt.Errorf("error: %w", err))
			return
		}

		if err := genServerRoutesFunction(gen, file, g, service, serviceInfo); err != nil {
			gen.Error(fmt.Errorf("error: %w", err))
			return
		}

		if err := genRoutesHandler(gen, file, g, service, serviceInfo); err != nil {
			gen.Error(fmt.Errorf("error: %w", err))
			return
		}

		if err := genClientWrapper(gen, file, g, service); err != nil {
			gen.Error(fmt.Errorf("error: %w", err))
			return
		}

		if err := genServerWrapper(gen, file, g, service); err != nil {
			gen.Error(fmt.Errorf("error: %w", err))
			return
		}
	}
}

func genClientWrapper(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) error {
	serverName := serverName(service)
	clientName := clientName(service)
	wrapperName := clientWrapperName(service)
	g.P("type ", wrapperName, " struct {")
	if *requireUnimplemented {
		g.P("Unimplemented", serverName)
	}
	g.P("cli ", clientName)
	g.P("options *", gorsPackage.Ident("Options"))
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			// Unary RPC method
			g.P("func (wrapper *", wrapperName, ") ", method.GoName, "(ctx ", contextPackage.Ident("Context"), ", request *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
			g.P("var headerMD, trailerMD ", metadataPackage.Ident("MD"))
			g.P("resp, err := wrapper.cli.", method.GoName, "(ctx, request, ", grpcPackage.Ident("Header"), "(&headerMD), ", grpcPackage.Ident("Trailer"), "(&trailerMD))")
			g.P(gorsPackage.Ident("AddGRPCMetadata"), "(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)")
			g.P("return resp, err")
			g.P("}")
			g.P()
		} else {
			// Streaming RPC method
			continue
		}
	}
	return nil
}

func genServerWrapper(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) error {
	serverName := serverName(service)
	wrapperName := serverWrapperName(service)
	g.P("type ", wrapperName, " struct {")
	if *requireUnimplemented {
		g.P("Unimplemented", serverName)
	}
	g.P("srv ", serverName)
	g.P("options *", gorsPackage.Ident("Options"))
	g.P("}")
	g.P()
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			// Unary RPC method
			g.P("func (wrapper *", wrapperName, ") ", method.GoName, "(ctx ", contextPackage.Ident("Context"), ", request *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
			g.P("rpcMethodName := ", strconv.Quote(fullMethodName(service, method)))
			g.P("stream := ", gorsPackage.Ident("NewServerTransportStream"), "(rpcMethodName)")
			g.P("ctx = ", grpcPackage.Ident("NewContextWithServerTransportStream"), "(ctx, stream)")
			g.P("resp, err := wrapper.srv.", method.GoName, "(ctx, request)")
			g.P(gorsPackage.Ident("AddGRPCMetadata"), "(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)")
			g.P("return resp, err")
			g.P("}")
			g.P()
		} else {
			// Streaming RPC method
			continue
		}
	}
	return nil
}

func getServiceInfo(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) (*parser.ServiceInfo, error) {
	serviceInfo, err := parser.NewService(splitComment(service.Comments.Leading.String()))
	if err != nil {
		return nil, err
	}
	serviceInfo.SetServiceName(service.GoName)
	serviceInfo.SetFullName(string(service.Desc.FullName()))
	var routers []*parser.RouterInfo
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			// Unary RPC method
			router, err := parser.ParseRouter(splitComment(method.Comments.Leading.String()))
			if err != nil {
				return nil, err
			}
			router.SetMethodName(method.GoName)
			router.SetFullMethodName(fullMethodName(service, method))

			if stringx.IsBlank(router.HttpMethod) {
				router.HttpMethod = parser.POST
			}
			if stringx.IsBlank(router.Path) {
				router.Path = router.FullMethodName
				if *pathToLower {
					router.Path = strings.ToLower(router.Path)
				}
			}
			if slicex.IsEmpty(router.Bindings) {
				router.Bindings = []parser.Binding{parser.ProtoJSONBinding}
				router.BindingContentType = parser.JSONContentType
			}
			if stringx.IsBlank(router.Render) {
				router.Render = parser.ProtoJSONRender
				router.RenderContentType = parser.JSONContentType
			}
			router.HandlerName = handlerName(service, method)
			router.ProtoMethod = method
			routers = append(routers, router)
		} else {
			// Streaming RPC method
			continue
		}
	}
	serviceInfo.Routers = routers
	return serviceInfo, nil
}

func genRoutesHandler(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, serviceInfo *parser.ServiceInfo) error {
	serverName := serverName(service)
	for _, router := range serviceInfo.Routers {
		g.P("func ", router.HandlerName, "(wrapper ", serverName, ", options *", gorsPackage.Ident("Options"), ") func(c *", ginPackage.Ident("Context"), ") {")
		g.P("return func(c *", ginPackage.Ident("Context"), ") {")
		g.P("var rpcMethodName = ", strconv.Quote(router.FullMethodName))
		g.P("var ctx = ", gorsPackage.Ident("NewContext"), "(c, rpcMethodName)")
		g.P("var req *", router.ProtoMethod.Input.GoIdent)
		g.P("var resp *", router.ProtoMethod.Output.GoIdent)
		g.P("var err error")
		g.P("req = new(", router.ProtoMethod.Input.GoIdent, ")")

		err := printRequestBinding(gen, g, router)
		if err != nil {
			return err
		}

		g.P("if ctx, err = ", gorsPackage.Ident("NewGRPCContext"), "(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {")
		g.P(gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
		g.P("return")
		g.P("}")

		g.P("resp, err = wrapper.", router.MethodName, "(ctx, req)")

		g.P("if err != nil {")
		g.P(gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
		g.P("return")
		g.P("}")

		if err := printResponseRender(gen, g, router); err != nil {
			return err
		}

		g.P("}")
		g.P("}")
		g.P("")
	}

	return nil
}

func genClientRoutesFunction(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, serviceInfo *parser.ServiceInfo) error {
	clientName := clientName(service)
	funcName := clientRoutesFunctionName(service)
	g.P("func ", funcName, "(cli ", clientName, ", opts ...", gorsPackage.Ident("Option"), ") []", gorsPackage.Ident("Route"), " {")
	g.P("options := ", gorsPackage.Ident("NewOptions"), "(opts...)")
	g.P("if len(options.Tag) == 0 && !options.DisableDefaultTag {")
	g.P("options.Tag = ", strconv.Quote("json"))
	g.P("}")
	g.P("wrapper := &", clientWrapperName(service), "{cli: cli, options: options}")
	g.P("return []", gorsPackage.Ident("Route"), "{")
	for _, router := range serviceInfo.Routers {
		p := path.Join(serviceInfo.BasePath, router.Path)
		g.P(gorsPackage.Ident("NewRoute"), "(", httpPackage.Ident(router.HttpMethod.HttpMethod()), ",", strconv.Quote(p), ",", router.HandlerName, "(wrapper, options),", "),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func genServerRoutesFunction(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, serviceInfo *parser.ServiceInfo) error {
	serverName := serverName(service)
	funcName := serverRoutesFunctionName(service)
	g.P("func ", funcName, "(srv ", serverName, ", opts ...", gorsPackage.Ident("Option"), ") []", gorsPackage.Ident("Route"), " {")
	g.P("options := ", gorsPackage.Ident("NewOptions"), "(opts...)")
	g.P("if len(options.Tag) == 0 && !options.DisableDefaultTag {")
	g.P("options.Tag = ", strconv.Quote("json"))
	g.P("}")
	g.P("wrapper := &", serverWrapperName(service), "{srv: srv, options: options}")
	g.P("return []", gorsPackage.Ident("Route"), "{")
	for _, router := range serviceInfo.Routers {
		p := path.Join(serviceInfo.BasePath, router.Path)
		g.P(gorsPackage.Ident("NewRoute"), "(", httpPackage.Ident(router.HttpMethod.HttpMethod()), ",", strconv.Quote(p), ",", router.HandlerName, "(wrapper, options),", "),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func printRequestBinding(gen *protogen.Plugin, g *protogen.GeneratedFile, router *parser.RouterInfo) error {

	g.P("if err = ", gorsPackage.Ident("RequestBind"), "(")
	g.P("ctx, req, options.Tag,")
	for _, binding := range router.Bindings {
		g.P(gorsPackage.Ident(strings.TrimPrefix(string(binding), "@")), ",")
	}
	g.P("); err != nil {")
	g.P(gorsPackage.Ident("ErrorRender"), "(ctx, err, options.ErrorHandler, options.ResponseWrapper)")
	g.P("return")
	g.P("}")
	return nil
}

func printResponseRender(gen *protogen.Plugin, g *protogen.GeneratedFile, router *parser.RouterInfo) error {
	renders := []parser.Render{
		parser.JSONRender, parser.IndentedJSONRender, parser.SecureJSONRender,
		parser.PureJSONRender, parser.AsciiJSONRender, parser.ProtoJSONRender,
		parser.ProtoBufRender, parser.CustomRender, parser.XMLRender,
		parser.YAMLRender, parser.TOMLRender, parser.MsgPackRender,
	}
	if !slices.Contains(renders, router.Render) {
		return fmt.Errorf("%s, %s is not supported", router.FullMethodName, router.Render)
	}

	renderName := strings.TrimPrefix(router.Render.String(), "@")
	renderArg := ""
	if router.Render == parser.ProtoJSONRender {
		renderArg = "(options.ProtoJSONMarshalOptions)"
	}
	g.P(gorsPackage.Ident("ResponseRender"),
		"(ctx, ", gorsPackage.Ident("StatusCode"), "(ctx), resp,",
		strconv.Quote(router.RenderContentType), ",", gorsPackage.Ident(renderName), renderArg,
		", options.ResponseWrapper)")

	return nil
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

func clientWrapperName(service *protogen.Service) string {
	return "_" + clientName(service) + "Wrapper"
}

func clientName(service *protogen.Service) string {
	return service.GoName + "Client"
}

func serverWrapperName(service *protogen.Service) string {
	return "_" + serverName(service) + "Wrapper"
}

func serverName(service *protogen.Service) string {
	return service.GoName + "Server"
}

func routesFunctionName(service *protogen.Service) string {
	return "_" + service.GoName + "Routes"
}

func clientRoutesFunctionName(service *protogen.Service) string {
	return clientName(service) + "Routes"
}

func serverRoutesFunctionName(service *protogen.Service) string {
	return serverName(service) + "Routes"
}

func fullMethodName(service *protogen.Service, method *protogen.Method) string {
	return fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())
}

func handlerName(service *protogen.Service, method *protogen.Method) string {
	return fmt.Sprintf("_%s_%s_GORS_Handler", service.Desc.Name(), method.Desc.Name())
}