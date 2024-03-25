package parser

import (
	"fmt"
	"github.com/go-leo/gox/slicex"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
	"path"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func LoadPkg(args []string) (*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedExportFile |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedTypesSizes,
	}
	pkgs, err := packages.Load(cfg, args...)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("error: %d packages found", len(pkgs))
	}
	return pkgs[0], nil
}

func Inspect(pkg *packages.Package, serviceName string) (*ast.File, *ast.GenDecl, *ast.TypeSpec, *ast.InterfaceType, []*ast.Field) {
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
				if typeSpec.Name.Name != serviceName {
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

func ExtractRouterInfo(method *ast.Field, methodName *ast.Ident) *RouterInfo {
	if method.Doc == nil {
		return NewRouter(methodName.String(), nil)
	}
	comments := slicex.Map[[]*ast.Comment, []string](
		method.Doc.List,
		func(i int, e1 *ast.Comment) string { return e1.Text },
	)
	return NewRouter(methodName.String(), comments)
}

func ExtractGoImports(serviceFile *ast.File) map[string]*GoImport {
	goImports := make(map[string]*GoImport)
	for _, importSpec := range serviceFile.Imports {
		importPath, err := strconv.Unquote(importSpec.Path.Value)
		if err != nil {
			log.Panicf("warning: unquote error: %s", err)
		}
		item := &GoImport{
			ImportPath: importPath,
		}
		if importSpec.Name != nil {
			item.PackageName = importSpec.Name.Name
		} else {
			item.PackageName = CleanPackageName(path.Base(importPath))
		}
		goImports[item.ImportPath] = item
	}
	return goImports
}

func InitServiceInfo(name string, serviceDecl *ast.GenDecl) *ServiceInfo {
	if serviceDecl == nil || serviceDecl.Doc == nil {
		return &ServiceInfo{Name: name}
	}
	var comments []string
	for _, comment := range serviceDecl.Doc.List {
		comments = append(comments, comment.Text)
	}
	return NewService(name, comments)
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

type GoImport struct {
	PackageName string
	ImportPath  string
	Enable      bool
}

type ObjectArgs struct {
	Name         string
	GoImportPath GoImportPath
	StarExpr     *ast.StarExpr
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

func CheckParams(rpcType *ast.FuncType, methodName *ast.Ident, imports map[string]*GoImport) (*Param, error) {
	if rpcType.Params == nil {
		return nil, fmt.Errorf("error: func %s params is empty", methodName)
	}
	if len(rpcType.Params.List) != 2 {
		return nil, fmt.Errorf("error: func %s params count is not equal 2", methodName)
	}
	// param1
	if err := CheckParam1MustBeContext(rpcType, methodName); err != nil {
		return nil, err
	}
	// param2
	param2, err := CheckAndGetParam2(rpcType, methodName, imports)
	if err != nil {
		return nil, err
	}
	return param2, nil
}

func CheckParam1MustBeContext(rpcType *ast.FuncType, methodName *ast.Ident) error {
	param1 := rpcType.Params.List[0]
	param0SelectorExpr, ok := param1.Type.(*ast.SelectorExpr)
	if !ok {
		return fmt.Errorf("error: func %s 1th param is not context.Context", methodName)
	}
	if param0SelectorExpr.Sel.Name != "Context" {
		return fmt.Errorf("error: func %s 1th param is not context.Context", methodName)
	}
	param0SelectorExprX, ok := param0SelectorExpr.X.(*ast.Ident)
	if !ok {
		return fmt.Errorf("error: func %s 1th param is not context.Context", methodName)
	}
	if param0SelectorExprX.Name != "context" {
		return fmt.Errorf("error: func %s 1th param is not context.Context", methodName)
	}
	return nil
}

func CheckAndGetParam2(rpcType *ast.FuncType, methodName *ast.Ident, imports map[string]*GoImport) (*Param, error) {
	param2 := rpcType.Params.List[1]
	errorTemplate := "error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}"
	switch p2 := param2.Type.(type) {
	case *ast.ArrayType:
		// []byte type
		ident, ok := p2.Elt.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf(errorTemplate, methodName.String())
		}
		if ident.Name != "byte" {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		return &Param{Bytes: true}, nil
	case *ast.Ident:
		// string type
		if p2.Name != "string" {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		return &Param{String: true}, nil
	case *ast.StarExpr:
		// *struct{} type
		switch x := p2.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &Param{ObjectArgs: &ObjectArgs{Name: name, StarExpr: p2}}, nil
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf(errorTemplate, methodName)
			}
			for importPath, goImport := range imports {
				if goImport.PackageName == ident.Name {
					return &Param{ObjectArgs: &ObjectArgs{Name: x.Sel.Name, GoImportPath: GoImportPath(importPath), StarExpr: p2}}, nil
				}
			}
			return nil, fmt.Errorf(errorTemplate, methodName)
		default:
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
	case *ast.SelectorExpr:
		// io.Reader type
		if p2.Sel == nil {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		if p2.Sel.Name != "Reader" {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		ident, ok := p2.X.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		ioImport, ok := imports["io"]
		if !ok {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		if ioImport.PackageName != ident.Name {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		return &Param{Reader: true}, nil
	default:
		return nil, fmt.Errorf(errorTemplate, methodName)
	}
}

func CheckResults(rpcType *ast.FuncType, methodName *ast.Ident, imports map[string]*GoImport) (*Result, error) {
	if rpcType.Results == nil {
		return nil, fmt.Errorf("error: func %s results is empty", methodName)
	}
	if len(rpcType.Results.List) != 2 {
		return nil, fmt.Errorf("error: func %s results count is not equal 2", methodName)
	}
	// result2
	if err := CheckResult2MustBeError(rpcType, methodName); err != nil {
		return nil, err
	}
	// result1
	result1, err := CheckAndGetResult1(rpcType, methodName, imports)
	if err != nil {
		return nil, err
	}
	return result1, nil
}

func CheckResult2MustBeError(rpcType *ast.FuncType, methodName *ast.Ident) error {
	result2 := rpcType.Results.List[1]
	result2Ident, ok := result2.Type.(*ast.Ident)
	if !ok {
		return fmt.Errorf("error: func %s 2th result is not error", methodName)
	}
	if result2Ident.Name != "error" {
		return fmt.Errorf("error: func %s 2th result is not error", methodName)
	}
	return nil
}

func CheckAndGetResult1(rpcType *ast.FuncType, methodName *ast.Ident, imports map[string]*GoImport) (*Result, error) {
	result1 := rpcType.Results.List[0]
	errorTemplate := "error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}"
	switch r1 := result1.Type.(type) {
	case *ast.ArrayType:
		// []byte type
		ident, ok := r1.Elt.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		if ident.Name != "byte" {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		return &Result{Bytes: true}, nil
	case *ast.Ident:
		// string type
		if r1.Name != "string" {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		return &Result{String: true}, nil
	case *ast.StarExpr:
		// *struct{} type
		switch x := r1.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &Result{ObjectArgs: &ObjectArgs{Name: name, StarExpr: r1}}, nil
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf(errorTemplate, methodName)
			}
			for importPath, goImport := range imports {
				if goImport.PackageName == ident.Name {
					return &Result{ObjectArgs: &ObjectArgs{Name: x.Sel.Name, GoImportPath: GoImportPath(importPath), StarExpr: r1}}, nil
				}
			}
			return nil, fmt.Errorf(errorTemplate, methodName)
		default:
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
	case *ast.SelectorExpr:
		// io.Reader type
		if r1.Sel == nil {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		if r1.Sel.Name != "Reader" {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		ident, ok := r1.X.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		ioImport, ok := imports["io"]
		if !ok {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		if ioImport.PackageName != ident.Name {
			return nil, fmt.Errorf(errorTemplate, methodName)
		}
		return &Result{Reader: true}, nil
	default:
		return nil, fmt.Errorf(errorTemplate, methodName)
	}
}
