package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/go-leo/gors/internal/pkg/annotation"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/gox/stringx"
	"go/ast"
	"go/format"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	serviceName = flag.String("service", "", "service interface Name; must be set")
	pathToLower = flag.Bool("path_to_lower", false, "make path to lower case")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of gors:\n")
	fmt.Fprintf(os.Stderr, "\tgors -service S\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
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
		pkgName:          pkg.Name,
		imports:          imports,
		srvName:          *serviceName,
		usedPackageNames: make(map[string]bool),
		routerInfos:      nil,
	}

	var basePath string
	if serviceDecl != nil && serviceSpec != nil && serviceType != nil && len(serviceMethods) > 0 {
		// find basePath
		basePath = extractBasePath(serviceDecl)
		// generate router by method comment
		for _, method := range serviceMethods {
			if slicex.IsEmpty(method.Names) {
				continue
			}
			methodName := method.Names[0]
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

			// results
			g.checkResults(rpcType, methodName)
			// result2
			g.checkResult2MustBeError(rpcType, methodName)
			// result1
			result1 := g.checkAndGetResult1(rpcType, methodName)

			fmName := fmt.Sprintf("/%s.%s/%s", g.pkgName, g.srvName, methodName.String())
			var routerInfo *annotation.RouterInfo
			if method.Doc == nil {
				routerInfo = annotation.NewRouter(methodName.String(), fmName, basePath, nil)
			} else {
				comments := slicex.Map[[]*ast.Comment, []string](
					method.Doc.List,
					func(i int, e1 *ast.Comment) string { return e1.Text },
				)
				routerInfo = annotation.NewRouter(methodName.String(), fmName, basePath, comments)
			}
			routerInfo.Param2 = param2
			routerInfo.Result1 = result1

			if stringx.IsBlank(routerInfo.HttpMethod) {
				routerInfo.HttpMethod = annotation.GET
			}
			if stringx.IsBlank(routerInfo.Path) {
				routerInfo.Path = routerInfo.FullMethodName
				if *pathToLower {
					routerInfo.Path = strings.ToLower(routerInfo.Path)
				}
			}
			defaultBindingName(routerInfo)
			defaultRenderName(routerInfo)
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

func defaultBindingName(info *annotation.RouterInfo) {
	if info.Param2.Bytes {
		info.Bindings = nil
	} else if info.Param2.String {
		info.Bindings = nil
	} else if info.Param2.Reader {
		info.Bindings = nil
	} else if objectArgs := info.Param2.ObjectArgs; objectArgs != nil {
		if slicex.IsEmpty(info.Bindings) {
			info.Bindings = []string{
				annotation.UriBinding,
				annotation.HeaderBinding,
				annotation.QueryBinding,
			}
		}
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.FullMethodName)
	}
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

func extractBasePath(serviceDecl *ast.GenDecl) string {
	if serviceDecl == nil || serviceDecl.Doc == nil {
		return ""
	}
	var comments []string
	for _, comment := range serviceDecl.Doc.List {
		comments = append(comments, comment.Text)
	}
	return annotation.ExtractBasePath(comments)
}

func getGoImports(serviceFile *ast.File) map[string]*annotation.GoImport {
	goImports := make(map[string]*annotation.GoImport)
	for _, importSpec := range serviceFile.Imports {
		importPath, err := strconv.Unquote(importSpec.Path.Value)
		if err != nil {
			log.Panicf("warning: unquote error: %s", err)
		}
		item := &annotation.GoImport{
			ImportPath: importPath,
		}
		if importSpec.Name != nil {
			item.PackageName = importSpec.Name.Name
		} else {
			item.PackageName = annotation.CleanPackageName(path.Base(importPath))
		}
		goImports[item.ImportPath] = item
	}
	return goImports
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

func defaultRenderName(info *annotation.RouterInfo) {
	switch {
	case info.Result1.Bytes:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.BytesRender
		}
	case info.Result1.String:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.StringRender
		}
	case info.Result1.Reader:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.ReaderRender
		}
	case info.Result1.ObjectArgs != nil:
		if stringx.IsBlank(info.Render) {
			info.Render = annotation.JSONRender
			info.RenderContentType = annotation.JSONContentType
		}
	default:
		log.Fatalf("error: func %s 1th result is invalid, must be io.Reader or []byte or string or *struct{}", info.FullMethodName)
	}
}
