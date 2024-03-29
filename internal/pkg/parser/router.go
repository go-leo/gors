package parser

import (
	"bytes"
	"fmt"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/gox/stringx"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"go/ast"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"path"
	"strings"
)

type RouterInfo struct {
	HttpMethod     Method
	Description    string
	Path           string
	MethodName     string
	FullMethodName string

	BindingContentType string
	Bindings           []Binding

	RenderContentType string
	Render            Render

	HandlerName string
	ProtoMethod *protogen.Method
	FuncType    *ast.FuncType
	Param2      *Param
	Result1     *Result
}

func (router *RouterInfo) PathParams() []string {
	var params []string
	segs := strings.Split(router.Path, "/")
	for _, seg := range segs {
		seg = strings.TrimSpace(seg)
		if stringx.IsBlank(seg) {
			continue
		}
		if strings.HasPrefix(seg, ":") {
			params = append(params, strings.TrimPrefix(seg, ":"))
			continue
		}
		if strings.HasPrefix(seg, "*") {
			params = append(params, strings.TrimPrefix(seg, "*"))
			continue
		}
		continue
	}
	return params
}

func (router *RouterInfo) QueryParams() []string {
	var params []string
	return params
}

func (router *RouterInfo) HeaderParams() []string {
	var params []string
	return params
}

func (router *RouterInfo) FormParams() []string {
	var params []string
	return params
}

func (router *RouterInfo) FileParams() []string {
	var params []string
	return params
}

func (router *RouterInfo) SetHandlerName(serviceName string) {
	router.HandlerName = fmt.Sprintf("_%s_%s_Handler", serviceName, router.MethodName)
}

func (router *RouterInfo) SetFullMethodName(rpcMethodName string) {
	router.FullMethodName = rpcMethodName
}

func (router *RouterInfo) SetFuncType(rpcType *ast.FuncType) {
	router.FuncType = rpcType
}

func (router *RouterInfo) SetParam2(param *Param) {
	router.Param2 = param
}

func (router *RouterInfo) SetResult1(result *Result) {
	router.Result1 = result
}

func (router *RouterInfo) SetMethodName(name string) {
	router.MethodName = name
}

func (router *RouterInfo) DefaultHttpMethod() {
	if stringx.IsBlank(router.HttpMethod) {
		router.HttpMethod = GET
	}
}

func (router *RouterInfo) DefaultHttpPath(pathToLower bool) {
	if stringx.IsBlank(router.Path) {
		router.Path = router.FullMethodName
		if pathToLower {
			router.Path = strings.ToLower(router.Path)
		}
	}
}

func (router *RouterInfo) DefaultBindingName() {
	Param2 := router.Param2
	if Param2.Reader {
		if slicex.IsEmpty(router.Bindings) {
			router.Bindings = []Binding{ReaderBinding}
		}
	} else if Param2.Bytes {
		if slicex.IsEmpty(router.Bindings) {
			router.Bindings = []Binding{BytesBinding}
		}
	} else if Param2.String {
		if slicex.IsEmpty(router.Bindings) {
			router.Bindings = []Binding{StringBinding}
		}
	} else if objectArgs := Param2.ObjectArgs; objectArgs != nil {
		if slicex.IsEmpty(router.Bindings) {
			router.Bindings = []Binding{QueryBinding}
			router.BindingContentType = ""
		}
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", router.FullMethodName)
	}
}

func (router *RouterInfo) DefaultRenderName() {
	Result1 := router.Result1
	switch {
	case Result1.Bytes:
		if stringx.IsBlank(router.Render) {
			router.Render = BytesRender
		}
	case Result1.String:
		if stringx.IsBlank(router.Render) {
			router.Render = StringRender
		}
	case Result1.Reader:
		if stringx.IsBlank(router.Render) {
			router.Render = ReaderRender
		}
	case Result1.ObjectArgs != nil:
		if stringx.IsBlank(router.Render) {
			router.Render = JSONRender
			router.RenderContentType = JSONContentType
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be io.Reader or []byte or string or *struct{}", router.FullMethodName)
	}
}

func (router *RouterInfo) OperationDoc(method Method) (*spec.Operation, error) {
	if method != router.HttpMethod {
		return nil, nil
	}
	parametersDoc, err := router.ParametersDoc()
	if err != nil {
		return nil, err
	}
	return &spec.Operation{
		VendorExtensible: spec.VendorExtensible{},
		OperationProps: spec.OperationProps{
			Description:  router.Description,
			Consumes:     []string{router.BindingContentType},
			Produces:     []string{router.RenderContentType},
			Schemes:      []string{},
			Tags:         []string{},
			Summary:      router.MethodName,
			ExternalDocs: &spec.ExternalDocumentation{Description: "", URL: ""},
			ID:           router.FullMethodName,
			Deprecated:   false,
			Security:     []map[string][]string{},
			Parameters:   parametersDoc,
			Responses: &spec.Responses{
				VendorExtensible: spec.VendorExtensible{},
				ResponsesProps: spec.ResponsesProps{
					Default:             nil,
					StatusCodeResponses: map[int]spec.Response{},
				},
			},
		},
	}, nil
}

func (router *RouterInfo) ParametersDoc() ([]spec.Parameter, error) {
	switch {
	case router.Param2.Bytes:
		fallthrough
	case router.Param2.String:
		fallthrough
	case router.Param2.Reader:
		return []spec.Parameter{router.rawBodyParameters()}, nil
	case router.Param2.ObjectArgs != nil:
		p2 := router.Param2.ObjectArgs.StarExpr
		switch x := p2.X.(type) {
		case *ast.Ident:
			if router.MethodName == "AllRequest" {
				fmt.Println(router)
			}
			typeSpec, ok := x.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				return nil, errors.New("failed x.Obj.Decl to *ast.TypeSpec")
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				return nil, errors.New("failed typeSpec.Type to *ast.StructType")
			}
			var parameters []spec.Parameter
			for _, field := range structType.Fields.List {
				name := field.Names[0].Name
				var desc []string
				if field.Doc != nil {
					for _, comment := range field.Doc.List {
						desc = append(desc, strings.TrimSpace(strings.Trim(strings.TrimSpace(comment.Text), "//")))
					}
				}
				if field.Comment != nil {
					for _, comment := range field.Comment.List {
						desc = append(desc, strings.TrimSpace(strings.Trim(strings.TrimSpace(comment.Text), "//")))
					}
				}

				switch typIdent := field.Type.(type) {
				case *ast.Ident:
					typ := typIdent.Name
					var in string
					if field.Tag != nil {
						tagValue := field.Tag.Value
						seg := strings.Split(strings.Trim(tagValue, "`"), ":")
						switch seg[0] {
						case "uri":
							in = "path"
						case "header":
							in = "header"
						case "form":
							if router.HttpMethod.EqualsIgnoreCase(GET.String()) {
								in = "query"
							} else {
								in = "formData"
							}
						default:
							in = "body"
						}
						name = strings.Split(strings.Trim(seg[1], `"`), ",")[0]
					}

					parameter := spec.Parameter{
						Refable:           spec.Refable{},
						CommonValidations: spec.CommonValidations{},
						SimpleSchema: spec.SimpleSchema{
							Type: typ,
						},
						VendorExtensible: spec.VendorExtensible{},
						ParamProps: spec.ParamProps{
							Description:     "",
							Name:            name,
							In:              in,
							Required:        true,
							Schema:          nil,
							AllowEmptyValue: false,
						},
					}
					parameters = append(parameters, parameter)
				default:
					return parameters, nil
					return nil, fmt.Errorf("unkown type %T", field.Type)
				}
			}
			_ = x
			return parameters, nil
		case *ast.SelectorExpr:
			var parameters []spec.Parameter
			_ = x
			return parameters, nil
		default:
			return nil, ErrParamType
		}
	}
	var parameters []spec.Parameter
	//for _, binding := range router.Bindings {
	//	switch binding {
	//	case ReaderBinding:
	//		parameters = append(parameters, router.rawBodyParameters())
	//	case BytesBinding:
	//		parameters = append(parameters, router.rawBodyParameters())
	//	case StringBinding:
	//		parameters = append(parameters, router.rawBodyParameters())
	//	case UriBinding:
	//		uriParameters, err := router.uriParameters()
	//		if err != nil {
	//			return nil, err
	//		}
	//		parameters = append(parameters, uriParameters...)
	//	case QueryBinding:
	//	case HeaderBinding:
	//	case JSONBinding:
	//	case XMLBinding:
	//	case FormBinding:
	//	case FormPostBinding:
	//	case FormMultipartBinding:
	//	case ProtoBufBinding:
	//	case MsgPackBinding:
	//	case YAMLBinding:
	//	case TOMLBinding:
	//	case ProtoJSONBinding:
	//	case CustomBinding:
	//	}
	//}
	return parameters, nil
}

func bbb() {
	//cfg := &packages.Config{
	//	Mode: packages.NeedName |
	//		packages.NeedFiles |
	//		packages.NeedCompiledGoFiles |
	//		packages.NeedImports |
	//		packages.NeedDeps |
	//		packages.NeedExportFile |
	//		packages.NeedTypes |
	//		packages.NeedSyntax |
	//		packages.NeedTypesInfo |
	//		packages.NeedTypesSizes |
	//		packages.NeedModule,
	//}
	//pkgs, err := packages.Load(cfg, goImport.PackageName)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(pkgs)
}

func (router *RouterInfo) rawBodyParameters() spec.Parameter {
	return spec.Parameter{
		ParamProps: spec.ParamProps{
			In:       "body",
			Required: true,
		},
	}
}

func (router *RouterInfo) uriParameters() ([]spec.Parameter, error) {
	wildcards, err := FindWildcards(router.Path)
	if err != nil {
		return nil, err
	}
	var r []spec.Parameter
	for _, wildcard := range wildcards {
		name := wildcard[1:]
		r = append(r, spec.Parameter{
			ParamProps: spec.ParamProps{
				Name:     name,
				In:       "path",
				Required: true,
			},
		})
	}
	return r, nil
}

func FindWildcards(path string) ([]string, error) {
	fullPath := path
	var wildcards []string
	for {
		// Find prefix until first wildcard
		wildcard, i, valid := FindWildcard(path)
		if i < 0 { // No wildcard found
			break
		}
		// The wildcard name must only contain one ':' or '*' character
		if !valid {
			return nil, errors.New("only one wildcard per path segment is allowed, has: '" + wildcard + "' in path '" + fullPath + "'")
		}
		// check if the wildcard has a name
		if len(wildcard) < 2 {
			return nil, errors.New("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}
		if wildcard[0] == ':' { // param
			if i > 0 {
				path = path[i:]
			}
			wildcards = append(wildcards, wildcard)
			// if the path doesn't end with the wildcard, then there
			// will be another subpath starting with '/'
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]
				continue
			}
			break
		}
	}
	return wildcards, nil
}

// FindWildcard Search for a wildcard segment and check the name for invalid characters.
// Returns -1 as index, if no wildcard was found.
func FindWildcard(path string) (wildcard string, i int, valid bool) {
	// Find start
	for start, c := range []byte(path) {
		// A wildcard starts with ':' (param) or '*' (catch-all)
		if c != ':' && c != '*' {
			continue
		}

		// Find end and check for invalid characters
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			case '/':
				return path[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}
		return path[start:], start, valid
	}
	return "", -1, false
}

var ErrMultipleHttpMethod = fmt.Errorf("there are multiple methods")

func ParseRouter(comments []string) (*RouterInfo, error) {
	r := &RouterInfo{}
	desc := &bytes.Buffer{}
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		// 注释的开始必须以 @GORS 开头
		if seg[0] != GORS {
			_, _ = fmt.Fprint(desc, text, " ")
			continue
		}
		if err := parseRouterComment(r, seg); err != nil {
			return nil, err
		}
	}
	r.Description = desc.String()
	return r, nil
}

func parseRouterComment(r *RouterInfo, comment []string) error {
	for _, segment := range comment {
		segment = strings.TrimSpace(segment)
		if segment == "" {
			continue
		}
		if startSegment(segment) {
			continue
		}

		pathSeg, ok, err := pathSegment(segment)
		if err != nil {
			return err
		}
		if ok {
			r.Path = path.Join(r.Path, pathSeg)
			continue
		}

		methodSeg, ok := httpMethodSegment(segment)
		if ok {
			if stringx.IsNotBlank(r.HttpMethod) {
				return ErrMultipleHttpMethod
			}
			r.HttpMethod = methodSeg
			continue
		}

		bindingSeg, contentType, ok := bindingSegment(segment)
		if ok {
			r.Bindings = append(r.Bindings, bindingSeg)
			r.BindingContentType = contentType
			continue
		}

		renderSeg, contentType, ok := renderSegment(segment)
		if ok {
			r.Render = renderSeg
			r.RenderContentType = contentType
			continue
		}
	}
	return nil
}

func startSegment(s string) bool {
	switch {
	case strings.HasPrefix(s, GORS):
		return true
	}
	return false
}

func pathSegment(s string) (string, bool, error) {
	switch {
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(Path)):
		v, ok := ExtractValue(s, Path)
		if !ok {
			return "", false, ErrPathInvalid
		}
		return v, true, nil
	}
	return "", false, nil
}

func httpMethodSegment(s string) (Method, bool) {
	// method start
	switch {
	case GET.EqualsIgnoreCase(s):
		return GET, true
	case POST.EqualsIgnoreCase(s):
		return POST, true
	case PUT.EqualsIgnoreCase(s):
		return PUT, true
	case DELETE.EqualsIgnoreCase(s):
		return DELETE, true
	case PATCH.EqualsIgnoreCase(s):
		return PATCH, true
	case HEAD.EqualsIgnoreCase(s):
		return HEAD, true
	case CONNECT.EqualsIgnoreCase(s):
		return CONNECT, true
	case OPTIONS.EqualsIgnoreCase(s):
		return OPTIONS, true
	case TRACE.EqualsIgnoreCase(s):
		return TRACE, true
	default:
		return "", false
	}
}

func bindingSegment(s string) (Binding, string, bool) {
	switch {
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(string(ReaderBinding))):
		v, _ := ExtractValue(s, string(ReaderBinding))
		return ReaderBinding, v, true
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(string(BytesBinding))):
		v, _ := ExtractValue(s, string(BytesBinding))
		return BytesBinding, v, true
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(string(StringBinding))):
		v, _ := ExtractValue(s, string(StringBinding))
		return StringBinding, v, true
	case strings.ToUpper(s) == strings.ToUpper(string(UriBinding)):
		return UriBinding, "", true
	case strings.ToUpper(s) == strings.ToUpper(string(QueryBinding)):
		return QueryBinding, "", true
	case strings.ToUpper(s) == strings.ToUpper(string(HeaderBinding)):
		return HeaderBinding, "", true
	case strings.ToUpper(s) == strings.ToUpper(string(JSONBinding)):
		return JSONBinding, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(ProtoJSONBinding)):
		return ProtoJSONBinding, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(XMLBinding)):
		return XMLBinding, XMLContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(FormBinding)):
		return FormBinding, FormContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(FormPostBinding)):
		return FormPostBinding, FormContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(FormMultipartBinding)):
		return FormMultipartBinding, FormMultipartContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(ProtoBufBinding)):
		return ProtoBufBinding, ProtoBufContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(MsgPackBinding)):
		return MsgPackBinding, MsgPackContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(YAMLBinding)):
		return YAMLBinding, YAMLContentType, true
	case strings.ToUpper(s) == strings.ToUpper(string(TOMLBinding)):
		return TOMLBinding, TOMLContentType, true
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(string(CustomBinding))):
		v, _ := ExtractValue(s, string(CustomBinding))
		return CustomBinding, v, true
	default:
		return "", "", false
	}
}

func renderSegment(s string) (Render, string, bool) {
	switch {
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(ReaderRender.String())):
		v, _ := ExtractValue(s, ReaderRender.String())
		return ReaderRender, v, true
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(BytesRender.String())):
		v, _ := ExtractValue(s, BytesRender.String())
		return BytesRender, v, true
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(StringRender.String())):
		v, _ := ExtractValue(s, StringRender.String())
		return StringRender, v, true
	case strings.ToUpper(s) == strings.ToUpper(TextRender.String()):
		return TextRender, PlainContentType, true
	case strings.ToUpper(s) == strings.ToUpper(HTMLRender.String()):
		return HTMLRender, HTMLContentType, true
	case strings.ToUpper(s) == strings.ToUpper(RedirectRender.String()):
		return RedirectRender, "", true
	case strings.ToUpper(s) == strings.ToUpper(JSONRender.String()):
		return JSONRender, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(IndentedJSONRender.String()):
		return IndentedJSONRender, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(SecureJSONRender.String()):
		return SecureJSONRender, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(JSONPJSONRender.String()):
		return JSONPJSONRender, JSONPContentType, true
	case strings.ToUpper(s) == strings.ToUpper(PureJSONRender.String()):
		return PureJSONRender, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(AsciiJSONRender.String()):
		return AsciiJSONRender, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(ProtoJSONRender.String()):
		return ProtoJSONRender, JSONContentType, true
	case strings.ToUpper(s) == strings.ToUpper(XMLRender.String()):
		return XMLRender, XMLContentType, true
	case strings.ToUpper(s) == strings.ToUpper(YAMLRender.String()):
		return YAMLRender, YAMLContentType, true
	case strings.ToUpper(s) == strings.ToUpper(ProtoBufRender.String()):
		return ProtoBufRender, ProtoBufContentType, true
	case strings.ToUpper(s) == strings.ToUpper(MsgPackRender.String()):
		return MsgPackRender, MsgPackContentType, true
	case strings.ToUpper(s) == strings.ToUpper(TOMLRender.String()):
		return TOMLRender, TOMLContentType, true
	case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(CustomRender.String())):
		v, _ := ExtractValue(s, CustomRender.String())
		return CustomRender, v, true
	default:
		return "", "", false
	}
}
