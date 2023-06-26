package gors

import (
	"github.com/go-leo/gors/internal/pkg/annotation"
	"github.com/go-leo/gors/internal/pkg/httpmethod"
	"github.com/go-leo/gox/stringx"
	"go/token"
	"log"
	"path"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ExtractBasePath(comments []string) string {
	var basePath string
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		if seg[0] != annotation.GORS {
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := FindPath(s)
				if !ok {
					log.Fatalf("error: %s path invalid", s)
				}
				basePath = path.Join(basePath, v)
			case strings.HasPrefix(s, annotation.GORS):
			case "" == s:
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	return basePath
}

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

type RouterInfo struct {
	Method string
	Path   string

	UriBinding           bool
	QueryBinding         bool
	HeaderBinding        bool
	JSONBinding          bool
	XMLBinding           bool
	FormBinding          bool
	FormPostBinding      bool
	FormMultipartBinding bool
	ProtoBufBinding      bool
	MsgPackBinding       bool
	YAMLBinding          bool
	TOMLBinding          bool
	CustomBinding        bool

	RenderContentType string

	BytesRender    bool
	StringRender   bool
	TextRender     bool
	HTMLRender     bool
	ReaderRender   bool
	RedirectRender bool

	JSONRender         bool
	IndentedJSONRender bool
	SecureJSONRender   bool
	JsonpJSONRender    bool
	PureJSONRender     bool
	AsciiJSONRender    bool
	XMLRender          bool
	YAMLRender         bool
	ProtobufRender     bool
	MsgPackRender      bool
	TOMLRender         bool
	CustomRender       bool

	RpcMethodName string
	Param2        *Param
	Result1       *Result
}

func NewRouter(methodName string, basePath string, comments []string) *RouterInfo {
	var r *RouterInfo
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		// 注释的开始必须以 @GORS 开头
		if seg[0] != annotation.GORS {
			continue
		}
		if r == nil {
			r = &RouterInfo{}
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := FindPath(s)
				if !ok {
					log.Fatalf("error: rpcmethod %s, %s path invalid", methodName, s)
				}
				r.Path = path.Join(r.Path, v)

				// method start
			case strings.ToUpper(s) == annotation.GET:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.GetMethod
			case strings.ToUpper(s) == annotation.POST:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.PostMethod
			case strings.ToUpper(s) == annotation.PUT:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.PutMethod
			case strings.ToUpper(s) == annotation.DELETE:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.DeleteMethod
			case strings.ToUpper(s) == annotation.PATCH:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.PatchMethod
			case strings.ToUpper(s) == annotation.HEAD:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.HeadMethod
			case strings.ToUpper(s) == annotation.CONNECT:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.ConnectMethod
			case strings.ToUpper(s) == annotation.OPTIONS:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.OptionsMethod
			case strings.ToUpper(s) == annotation.TRACE:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.TraceMethod
				// method end

				// binding start
			case strings.ToUpper(s) == strings.ToUpper(annotation.UriBinding):
				r.UriBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.QueryBinding):
				r.QueryBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.HeaderBinding):
				r.HeaderBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONBinding):
				r.JSONBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLBinding):
				r.XMLBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormBinding):
				r.FormBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormPostBinding):
				r.FormPostBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormMultipartBinding):
				r.FormMultipartBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufBinding):
				r.ProtoBufBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackBinding):
				r.MsgPackBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLBinding):
				r.YAMLBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLBinding):
				r.TOMLBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.CustomBinding):
				r.CustomBinding = true
				// binding end

				// render start
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.BytesRender)):
				v, _ := findBytesRender(s)
				r.RenderContentType = v
				r.BytesRender = true
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.StringRender)):
				v, _ := findStringRender(s)
				r.RenderContentType = v
				r.StringRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.TextRender):
				r.RenderContentType = "text/plain; charset=utf-8"
				r.TextRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.HTMLRender):
				r.RenderContentType = "text/html; charset=utf-8"
				r.HTMLRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.RedirectRender):
				r.RedirectRender = true
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.ReaderRender)):
				v, _ := findReaderRender(s)
				r.RenderContentType = v
				r.ReaderRender = true

			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONRender):
				r.JSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.IndentedJSONRender):
				r.IndentedJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.SecureJSONRender):
				r.SecureJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.JsonpJSONRender):
				r.JsonpJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.PureJSONRender):
				r.PureJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.AsciiJSONRender):
				r.AsciiJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLRender):
				r.XMLRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLRender):
				r.YAMLRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufRender):
				r.ProtobufRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackRender):
				r.MsgPackRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLRender):
				r.TOMLRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.CustomRender):
				r.CustomRender = true
				// render end

			case strings.HasPrefix(s, annotation.GORS):
			case "" == s:
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	if r != nil {
		if stringx.IsBlank(r.Method) {
			log.Fatalf("error: rpcmethod %s, http method is empty", methodName)
		}
		r.Path = path.Join(basePath, r.Path)
	}
	return r
}

func FindPath(s string) (string, bool) {
	reg := regexp.MustCompile(`@Path\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func findStringRender(s string) (string, bool) {
	reg := regexp.MustCompile(`@StringRender\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func findBytesRender(s string) (string, bool) {
	reg := regexp.MustCompile(`@BytesRender\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func findReaderRender(s string) (string, bool) {
	reg := regexp.MustCompile(`@ReaderRender\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
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
