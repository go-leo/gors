package gors

import (
	"github.com/go-leo/gors/internal/pkg/annotation"
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
	Method            annotation.Method
	Path              string
	Bindings          []string
	RenderContentType string
	Render            string
	RpcMethodName     string
	Param2            *Param
	Result1           *Result
}

func NewRouter(methodName string, basePath string, comments []string) *RouterInfo {
	r := &RouterInfo{}
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		// 注释的开始必须以 @GORS 开头
		if seg[0] != annotation.GORS {
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {

			// path
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := FindPath(s)
				if !ok {
					log.Fatalf("error: rpcmethod %s, %s path invalid", methodName, s)
				}
				r.Path = path.Join(r.Path, v)

				// method start
			case annotation.GET.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.GET
			case annotation.POST.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.POST
			case annotation.PUT.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.PUT
			case annotation.DELETE.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.DELETE
			case annotation.PATCH.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.PATCH
			case annotation.HEAD.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.HEAD
			case annotation.CONNECT.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.CONNECT
			case annotation.OPTIONS.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.OPTIONS
			case annotation.TRACE.EqualsIgnoreCase(s):
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = annotation.TRACE
				// method end

				// binding start
			case strings.ToUpper(s) == strings.ToUpper(annotation.UriBinding):
				r.Bindings = append(r.Bindings, annotation.UriBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.QueryBinding):
				r.Bindings = append(r.Bindings, annotation.QueryBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.HeaderBinding):
				r.Bindings = append(r.Bindings, annotation.HeaderBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONBinding):
				r.Bindings = append(r.Bindings, annotation.JSONBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoJSONBinding):
				r.Bindings = append(r.Bindings, annotation.ProtoJSONBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLBinding):
				r.Bindings = append(r.Bindings, annotation.XMLBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormBinding):
				r.Bindings = append(r.Bindings, annotation.FormBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormPostBinding):
				r.Bindings = append(r.Bindings, annotation.FormPostBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormMultipartBinding):
				r.Bindings = append(r.Bindings, annotation.FormMultipartBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufBinding):
				r.Bindings = append(r.Bindings, annotation.ProtoBufBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackBinding):
				r.Bindings = append(r.Bindings, annotation.MsgPackBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLBinding):
				r.Bindings = append(r.Bindings, annotation.YAMLBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLBinding):
				r.Bindings = append(r.Bindings, annotation.TOMLBinding)
			case strings.ToUpper(s) == strings.ToUpper(annotation.CustomBinding):
				r.Bindings = append(r.Bindings, annotation.CustomBinding)
				// binding end

				// render start
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.BytesRender)):
				v, _ := findBytesRender(s)
				r.RenderContentType = v
				r.Render = annotation.BytesRender
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.StringRender)):
				v, _ := findStringRender(s)
				r.RenderContentType = v
				r.Render = annotation.StringRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.TextRender):
				r.Render = annotation.TextRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.HTMLRender):
				r.Render = annotation.HTMLRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.RedirectRender):
				r.Render = annotation.RedirectRender
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.ReaderRender)):
				v, _ := findReaderRender(s)
				r.RenderContentType = v
				r.Render = annotation.ReaderRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONRender):
				r.Render = annotation.JSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.IndentedJSONRender):
				r.Render = annotation.IndentedJSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.SecureJSONRender):
				r.Render = annotation.SecureJSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONPJSONRender):
				r.Render = annotation.JSONPJSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.PureJSONRender):
				r.Render = annotation.PureJSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.AsciiJSONRender):
				r.Render = annotation.AsciiJSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoJSONRender):
				r.Render = annotation.ProtoJSONRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLRender):
				r.Render = annotation.XMLRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLRender):
				r.Render = annotation.YAMLRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufRender):
				r.Render = annotation.ProtoBufRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackRender):
				r.Render = annotation.MsgPackRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLRender):
				r.Render = annotation.TOMLRender
			case strings.ToUpper(s) == strings.ToUpper(annotation.CustomRender):
				r.Render = annotation.CustomRender
				// render end

			case strings.HasPrefix(s, annotation.GORS):
			case "" == s:
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	//if stringx.IsBlank(r.Method) {
	//	r.Method = httpmethod.PostMethod
	//	log.Printf("rpcmethod %s, use default http method POST", methodName)
	//}
	r.Path = path.Join(basePath, r.Path)
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
