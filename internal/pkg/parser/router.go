package parser

import (
	"bytes"
	"fmt"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/gox/stringx"
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

func (routerInfo *RouterInfo) PathParams() []string {
	var params []string
	segs := strings.Split(routerInfo.Path, "/")
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

func (routerInfo *RouterInfo) QueryParams() []string {
	var params []string
	return params
}

func (routerInfo *RouterInfo) HeaderParams() []string {
	var params []string
	return params
}

func (routerInfo *RouterInfo) FormParams() []string {
	var params []string
	return params
}

func (routerInfo *RouterInfo) FileParams() []string {
	var params []string
	return params
}

func (routerInfo *RouterInfo) SetHandlerName(serviceName string) {
	routerInfo.HandlerName = fmt.Sprintf("_%s_%s_Handler", serviceName, routerInfo.MethodName)
}

func (routerInfo *RouterInfo) SetFullMethodName(rpcMethodName string) {
	routerInfo.FullMethodName = rpcMethodName
}

func (routerInfo *RouterInfo) SetFuncType(rpcType *ast.FuncType) {
	routerInfo.FuncType = rpcType
}

func (routerInfo *RouterInfo) SetParam2(param *Param) {
	routerInfo.Param2 = param
}

func (routerInfo *RouterInfo) SetResult1(result *Result) {
	routerInfo.Result1 = result
}

func (routerInfo *RouterInfo) SetMethodName(name string) {
	routerInfo.MethodName = name
}

func (routerInfo *RouterInfo) DefaultHttpMethod() {
	if stringx.IsBlank(routerInfo.HttpMethod) {
		routerInfo.HttpMethod = GET
	}
}

func (routerInfo *RouterInfo) DefaultHttpPath(pathToLower bool) {
	if stringx.IsBlank(routerInfo.Path) {
		routerInfo.Path = routerInfo.FullMethodName
		if pathToLower {
			routerInfo.Path = strings.ToLower(routerInfo.Path)
		}
	}
}

func (routerInfo *RouterInfo) DefaultBindingName() {
	Param2 := routerInfo.Param2
	if Param2.Reader {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []Binding{ReaderBinding}
		}
	} else if Param2.Bytes {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []Binding{BytesBinding}
		}
	} else if Param2.String {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []Binding{StringBinding}
		}
	} else if objectArgs := Param2.ObjectArgs; objectArgs != nil {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []Binding{QueryBinding}
			routerInfo.BindingContentType = ""
		}
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", routerInfo.FullMethodName)
	}
}

func (routerInfo *RouterInfo) DefaultRenderName() {
	Result1 := routerInfo.Result1
	switch {
	case Result1.Bytes:
		if stringx.IsBlank(routerInfo.Render) {
			routerInfo.Render = BytesRender
		}
	case Result1.String:
		if stringx.IsBlank(routerInfo.Render) {
			routerInfo.Render = StringRender
		}
	case Result1.Reader:
		if stringx.IsBlank(routerInfo.Render) {
			routerInfo.Render = ReaderRender
		}
	case Result1.ObjectArgs != nil:
		if stringx.IsBlank(routerInfo.Render) {
			routerInfo.Render = JSONRender
			routerInfo.RenderContentType = JSONContentType
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be io.Reader or []byte or string or *struct{}", routerInfo.FullMethodName)
	}
}

var (
	strColon = []byte(":")
	strStar  = []byte("*")
	strSlash = []byte("/")
)

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
			return "", false, fmt.Errorf("%s path invalid", s)
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
