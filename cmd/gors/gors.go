package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-leo/slicex"
	"github.com/go-leo/stringx"
	"golang.org/x/tools/go/packages"

	"github.com/go-leo/gors/internal/pkg/annotation"
	"github.com/go-leo/gors/internal/pkg/httpmethod"
)

var (
	serviceName = flag.String("service", "", "service interface name; must be set")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of gors:\n")
	fmt.Fprintf(os.Stderr, "\tgors -service S\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

type goImportPath string

func (p goImportPath) Ident(s string) *goIdent {
	importPath := string(p)
	return &goIdent{
		GoName: s,
		GoImport: &goImport{
			PackageName: cleanPackageName(path.Base(importPath)),
			ImportPath:  importPath,
		},
	}
}

type goIdent struct {
	GoImport *goImport
	GoName   string
}

func (x *goIdent) Qualify() string {
	if x.GoImport.ImportPath == "" {
		return x.GoName
	}
	return x.GoImport.PackageName + "." + x.GoName
}

type goImport struct {
	PackageName string
	ImportPath  string
	enable      bool
}

type objectArgs struct {
	name         string
	goImportPath goImportPath
}

type param struct {
	bytes      bool
	string     bool
	objectArgs *objectArgs
	reader     bool
}

type result struct {
	bytes      bool
	string     bool
	objectArgs *objectArgs
	reader     bool
}

type routerInfo struct {
	method string
	path   string

	uriBinding           bool
	queryBinding         bool
	headerBinding        bool
	jsonBinding          bool
	xmlBinding           bool
	formBinding          bool
	formPostBinding      bool
	formMultipartBinding bool
	protobufBinding      bool
	msgpackBinding       bool
	yamlBinding          bool
	tomlBinding          bool

	rpcMethodName string
	param2        *param
	result1       *result

	renderContentType string

	bytesRender    bool
	stringRender   bool
	textRender     bool
	htmlRender     bool
	readerRender   bool
	redirectRender bool

	jsonRender         bool
	indentedJSONRender bool
	secureJSONRender   bool
	jsonpJSONRender    bool
	pureJSONRender     bool
	asciiJSONRender    bool
	xmlRender          bool
	yamlRender         bool
	protobufRender     bool
	msgpackRender      bool
	tomlRender         bool
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("gors: ")
}

func main() {

	flag.Usage = Usage
	flag.Parse()

	// must set service names
	if len(*serviceName) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// load package information
	pkg := loadPkg(args)

	// inspect package
	serviceFile, serviceDecl, serviceSpec, serviceType, serviceMethods := inspect(pkg)
	if serviceFile == nil || serviceDecl == nil || serviceSpec == nil || serviceType == nil {
		log.Fatal("error: not found service")
	}

	imports := getGoImports(serviceFile)
	g := &generate{
		buf:              &bytes.Buffer{},
		headerBuf:        &bytes.Buffer{},
		importsBuf:       &bytes.Buffer{},
		functionBuf:      &bytes.Buffer{},
		header:           fmt.Sprintf(`// Code generated by "gors %s"; DO NOT EDIT.`, strings.Join(os.Args[1:], " ")),
		pkg:              fmt.Sprintf("package %s", pkg.Name),
		imports:          imports,
		srvName:          *serviceName,
		routerInfos:      nil,
		usedPackageNames: make(map[string]bool),
	}

	var basePath string
	if serviceDecl != nil && serviceSpec != nil && serviceType != nil && len(serviceMethods) > 0 {
		// find basePath\globalConsume\globalProduce
		basePath = getBaseInfo(serviceDecl)
		// generate router by method comment
		for _, method := range serviceMethods {
			if method.Doc == nil && len(method.Doc.List) <= 0 {
				continue
			}
			if slicex.IsEmpty(method.Names) {
				continue
			}
			methodName := method.Names[0]
			routerInfo := newRouter(methodName, basePath, method.Doc.List)
			if routerInfo == nil {
				continue
			}
			routerInfo.rpcMethodName = methodName.Name
			rpcType, ok := method.Type.(*ast.FuncType)
			if !ok {
				log.Fatalf("error: func %s not convert to *ast.FuncType", methodName)
			}

			// params
			g.checkParams(rpcType, methodName)
			// param1
			g.checkParam1MustBeContext(rpcType, methodName)
			// param2
			param2 := g.checkAndGetParam2(rpcType, methodName)
			routerInfo.param2 = param2

			// results
			g.checkResults(rpcType, methodName)
			// result2
			g.checkResult2MustBeError(rpcType, methodName)
			// result1
			result1 := g.checkAndGetResult1(rpcType, methodName)
			routerInfo.result1 = result1

			g.routerInfos = append(g.routerInfos, routerInfo)
		}
	}

	content := g.content()
	// Format the output.
	src, err := format.Source(content)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		src = content
	}

	// Write to file.
	outDir, err := detectOutputDir(pkg.GoFiles)
	if err != nil {
		log.Fatalf("error: detect output dir: %s", err)
	}
	outputPath := filepath.Join(outDir, fmt.Sprintf("%s_gors.go", strings.ToLower(*serviceName)))

	if err := os.WriteFile(outputPath, src, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
	log.Printf("%s.%s wrote %s", pkg.PkgPath, *serviceName, outputPath)
}

func getBaseInfo(serviceDecl *ast.GenDecl) string {
	var basePath string
	for _, comment := range serviceDecl.Doc.List {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment.Text), "//"))
		seg := strings.Split(text, " ")
		if seg[0] != annotation.GORS {
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := findPath(s)
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

func getGoImports(serviceFile *ast.File) map[string]*goImport {
	goImports := make(map[string]*goImport)
	for _, importSpec := range serviceFile.Imports {
		importPath, err := strconv.Unquote(importSpec.Path.Value)
		if err != nil {
			log.Panicf("warning: unquote error: %s", err)
		}
		item := &goImport{
			ImportPath: importPath,
		}
		if importSpec.Name != nil {
			item.PackageName = importSpec.Name.Name
		} else {
			item.PackageName = cleanPackageName(path.Base(importPath))
		}
		goImports[item.ImportPath] = item
	}
	return goImports
}

func loadPkg(args []string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps | packages.NeedExportFile | packages.NeedTypes |
			packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
	}
	pkgs, err := packages.Load(cfg, args...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	pkg := pkgs[0]
	return pkg
}

func inspect(pkg *packages.Package) (*ast.File, *ast.GenDecl, *ast.TypeSpec, *ast.InterfaceType, []*ast.Field) {
	var serviceFile *ast.File
	var serviceDecl *ast.GenDecl
	var serviceSpec *ast.TypeSpec
	var serviceType *ast.InterfaceType
	var serviceMethods []*ast.Field
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			if node == nil {
				return true
			}
			denDecl, ok := node.(*ast.GenDecl)
			if !ok {
				return true
			}
			if denDecl.Tok != token.TYPE {
				// We only care about type declarations.
				return true
			}
			for _, spec := range denDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}
				if typeSpec.Name.Name != *serviceName {
					// This is not the interface type we're looking for.
					continue
				}
				serviceFile = file
				serviceDecl = denDecl
				serviceSpec = typeSpec
				serviceType = interfaceType
				serviceMethods = interfaceType.Methods.List
				return false
			}
			return true
		})
	}
	return serviceFile, serviceDecl, serviceSpec, serviceType, serviceMethods
}

func newRouter(methodName *ast.Ident, basePath string, commentList []*ast.Comment) *routerInfo {
	var r *routerInfo
	for _, comment := range commentList {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment.Text), "//"))
		seg := strings.Split(text, " ")
		// ???????????????????????? @GORS ??????
		if seg[0] != annotation.GORS {
			continue
		}
		if r == nil {
			r = &routerInfo{}
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := findPath(s)
				if !ok {
					log.Fatalf("error: rpcmethod %s, %s path invalid", methodName.String(), s)
				}
				r.path = path.Join(r.path, v)

				// method start
			case strings.ToUpper(s) == annotation.GET:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.GetMethod
			case strings.ToUpper(s) == annotation.POST:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.PostMethod
			case strings.ToUpper(s) == annotation.PUT:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.PutMethod
			case strings.ToUpper(s) == annotation.DELETE:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.DeleteMethod
			case strings.ToUpper(s) == annotation.PATCH:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.PatchMethod
			case strings.ToUpper(s) == annotation.HEAD:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.HeadMethod
			case strings.ToUpper(s) == annotation.CONNECT:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.ConnectMethod
			case strings.ToUpper(s) == annotation.OPTIONS:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.OptionsMethod
			case strings.ToUpper(s) == annotation.TRACE:
				if stringx.IsNotBlank(r.method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName.String())
				}
				r.method = httpmethod.TraceMethod
				// method end

				// binding start
			case strings.ToUpper(s) == strings.ToUpper(annotation.UriBinding):
				r.uriBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.QueryBinding):
				r.queryBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.HeaderBinding):
				r.headerBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONBinding):
				r.jsonBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLBinding):
				r.xmlBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormBinding):
				r.formBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormPostBinding):
				r.formPostBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormMultipartBinding):
				r.formMultipartBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufBinding):
				r.protobufBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackBinding):
				r.msgpackBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLBinding):
				r.yamlBinding = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLBinding):
				r.tomlBinding = true
				// binding end

				// render start
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.BytesRender)):
				v, _ := findBytesRender(s)
				r.renderContentType = v
				r.bytesRender = true
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.StringRender)):
				v, _ := findStringRender(s)
				r.renderContentType = v
				r.stringRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.TextRender):
				r.renderContentType = "text/plain; charset=utf-8"
				r.textRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.HTMLRender):
				r.renderContentType = "text/html; charset=utf-8"
				r.htmlRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.RedirectRender):
				r.redirectRender = true
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.ReaderRender)):
				v, _ := findReaderRender(s)
				r.renderContentType = v
				r.readerRender = true

			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONRender):
				r.jsonRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.IndentedJSONRender):
				r.indentedJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.SecureJSONRender):
				r.secureJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.JsonpJSONRender):
				r.jsonpJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.PureJSONRender):
				r.pureJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.AsciiJSONRender):
				r.asciiJSONRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLRender):
				r.xmlRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLRender):
				r.yamlRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufRender):
				r.protobufRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackRender):
				r.msgpackRender = true
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLRender):
				r.tomlRender = true
				// render end

			case strings.HasPrefix(s, annotation.GORS):
			case "" == s:
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	if r != nil {
		if stringx.IsBlank(r.method) {
			log.Fatalf("error: rpcmethod %s, http method is empty", methodName.String())
		}
		r.path = path.Join(basePath, r.path)
	}
	return r
}

func findPath(s string) (string, bool) {
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

func cleanPackageName(name string) string {
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

func detectOutputDir(paths []string) (string, error) {
	if len(paths) == 0 {
		return "", errors.New("no files to derive output directory from")
	}
	dir := filepath.Dir(paths[0])
	for _, p := range paths[1:] {
		if dir2 := filepath.Dir(p); dir2 != dir {
			return "", fmt.Errorf("found conflicting directories %q and %q", dir, dir2)
		}
	}
	return dir, nil
}
