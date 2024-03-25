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
	HttpMethod         Method
	Description        string
	Path               string
	MethodName         string
	FullMethodName     string
	BindingContentType string
	Bindings           []string
	RenderContentType  string
	Render             string
	HandlerName        string
	ProtoMethod        *protogen.Method
	FuncType           *ast.FuncType
	Param2             *Param
	Result1            *Result
}

func (info *RouterInfo) PathParams() []string {
	var params []string
	segs := strings.Split(info.Path, "/")
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

func (info *RouterInfo) QueryParams() []string {
	var params []string
	return params
}

func (info *RouterInfo) HeaderParams() []string {
	var params []string
	return params
}

func (info *RouterInfo) FormParams() []string {
	var params []string
	return params
}

func (info *RouterInfo) FileParams() []string {
	var params []string
	return params
}

func (info *RouterInfo) SetHandlerName(serviceName string) {
	info.HandlerName = fmt.Sprintf("_%s_%s_Handler", serviceName, info.MethodName)
}

func (info *RouterInfo) SetFullMethodName(rpcMethodName string) {
	info.FullMethodName = rpcMethodName
}

func (info *RouterInfo) SetFuncType(rpcType *ast.FuncType) {
	info.FuncType = rpcType
}

func (info *RouterInfo) SetParam2(param *Param) {
	info.Param2 = param
}

func (info *RouterInfo) SetResult1(result *Result) {
	info.Result1 = result
}

func (info *RouterInfo) SetMethodName(name string) {
	info.MethodName = name
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
			routerInfo.Bindings = []string{ReaderBinding}
		}
	} else if Param2.Bytes {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []string{BytesBinding}
		}
	} else if Param2.String {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []string{StringBinding}
		}
	} else if objectArgs := Param2.ObjectArgs; objectArgs != nil {
		if slicex.IsEmpty(routerInfo.Bindings) {
			routerInfo.Bindings = []string{QueryBinding}
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
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(s, GORS):
				continue
			case "" == s:
				continue

				// path
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(Path)):
				v, ok := ExtractValue(s, Path)
				if !ok {
					return nil, fmt.Errorf("%s path invalid", s)
				}
				r.Path = path.Join(r.Path, v)

				// method start
			case GET.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = GET
			case POST.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = POST
			case PUT.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = PUT
			case DELETE.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = DELETE
			case PATCH.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = PATCH
			case HEAD.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = HEAD
			case CONNECT.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = CONNECT
			case OPTIONS.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = OPTIONS
			case TRACE.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.HttpMethod) {
					return nil, ErrMultipleHttpMethod
				}
				r.HttpMethod = TRACE
				// method end

				// binding start
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(ReaderBinding)):
				v, _ := ExtractValue(s, ReaderBinding)
				r.BindingContentType = v
				r.Bindings = append(r.Bindings, ReaderBinding)
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(BytesBinding)):
				v, _ := ExtractValue(s, BytesBinding)
				r.BindingContentType = v
				r.Bindings = append(r.Bindings, BytesBinding)
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(StringBinding)):
				v, _ := ExtractValue(s, StringBinding)
				r.BindingContentType = v
				r.Bindings = append(r.Bindings, StringBinding)
			case strings.ToUpper(s) == strings.ToUpper(UriBinding):
				r.Bindings = append(r.Bindings, UriBinding)
			case strings.ToUpper(s) == strings.ToUpper(QueryBinding):
				r.Bindings = append(r.Bindings, QueryBinding)
			case strings.ToUpper(s) == strings.ToUpper(HeaderBinding):
				r.Bindings = append(r.Bindings, HeaderBinding)
			case strings.ToUpper(s) == strings.ToUpper(JSONBinding):
				r.Bindings = append(r.Bindings, JSONBinding)
				r.BindingContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(ProtoJSONBinding):
				r.Bindings = append(r.Bindings, ProtoJSONBinding)
				r.BindingContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(XMLBinding):
				r.Bindings = append(r.Bindings, XMLBinding)
				r.BindingContentType = XMLContentType
			case strings.ToUpper(s) == strings.ToUpper(FormBinding):
				r.Bindings = append(r.Bindings, FormBinding)
				r.BindingContentType = FormContentType
			case strings.ToUpper(s) == strings.ToUpper(FormPostBinding):
				r.Bindings = append(r.Bindings, FormPostBinding)
				r.BindingContentType = FormContentType
			case strings.ToUpper(s) == strings.ToUpper(FormMultipartBinding):
				r.Bindings = append(r.Bindings, FormMultipartBinding)
				r.BindingContentType = FormMultipartContentType
			case strings.ToUpper(s) == strings.ToUpper(ProtoBufBinding):
				r.Bindings = append(r.Bindings, ProtoBufBinding)
				r.BindingContentType = ProtoBufContentType
			case strings.ToUpper(s) == strings.ToUpper(MsgPackBinding):
				r.Bindings = append(r.Bindings, MsgPackBinding)
				r.BindingContentType = MsgPackContentType
			case strings.ToUpper(s) == strings.ToUpper(YAMLBinding):
				r.Bindings = append(r.Bindings, YAMLBinding)
				r.BindingContentType = YAMLContentType
			case strings.ToUpper(s) == strings.ToUpper(TOMLBinding):
				r.Bindings = append(r.Bindings, TOMLBinding)
				r.BindingContentType = TOMLContentType
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(CustomBinding)):
				v, _ := ExtractValue(s, CustomBinding)
				r.BindingContentType = v
				r.Bindings = append(r.Bindings, CustomBinding)
				// binding end

				// render start
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(ReaderRender)):
				v, _ := ExtractValue(s, ReaderRender)
				r.RenderContentType = v
				r.Render = ReaderRender
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(BytesRender)):
				v, _ := ExtractValue(s, BytesRender)
				r.RenderContentType = v
				r.Render = BytesRender
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(StringRender)):
				v, _ := ExtractValue(s, StringRender)
				r.RenderContentType = v
				r.Render = StringRender
			case strings.ToUpper(s) == strings.ToUpper(TextRender):
				r.Render = TextRender
				r.RenderContentType = PlainContentType
			case strings.ToUpper(s) == strings.ToUpper(HTMLRender):
				r.Render = HTMLRender
				r.RenderContentType = HTMLContentType
			case strings.ToUpper(s) == strings.ToUpper(RedirectRender):
				r.Render = RedirectRender
			case strings.ToUpper(s) == strings.ToUpper(JSONRender):
				r.Render = JSONRender
				r.RenderContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(IndentedJSONRender):
				r.Render = IndentedJSONRender
				r.RenderContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(SecureJSONRender):
				r.Render = SecureJSONRender
				r.RenderContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(JSONPJSONRender):
				r.Render = JSONPJSONRender
				r.RenderContentType = JSONPContentType
			case strings.ToUpper(s) == strings.ToUpper(PureJSONRender):
				r.Render = PureJSONRender
				r.RenderContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(AsciiJSONRender):
				r.Render = AsciiJSONRender
				r.RenderContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(ProtoJSONRender):
				r.Render = ProtoJSONRender
				r.RenderContentType = JSONContentType
			case strings.ToUpper(s) == strings.ToUpper(XMLRender):
				r.Render = XMLRender
				r.RenderContentType = XMLContentType
			case strings.ToUpper(s) == strings.ToUpper(YAMLRender):
				r.Render = YAMLRender
				r.RenderContentType = YAMLContentType
			case strings.ToUpper(s) == strings.ToUpper(ProtoBufRender):
				r.Render = ProtoBufRender
				r.RenderContentType = ProtoBufContentType
			case strings.ToUpper(s) == strings.ToUpper(MsgPackRender):
				r.Render = MsgPackRender
				r.RenderContentType = MsgPackContentType
			case strings.ToUpper(s) == strings.ToUpper(TOMLRender):
				r.Render = TOMLRender
				r.RenderContentType = TOMLContentType
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(CustomRender)):
				v, _ := ExtractValue(s, CustomRender)
				r.RenderContentType = v
				r.Render = CustomRender
				// render end

			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	r.Description = desc.String()
	return r, nil
}
