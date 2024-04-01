package parser

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"
)

func ParseService(args []string, serviceName string, pathToLower bool) (*ServiceInfo, error) {
	// load package information
	pkg, err := LoadPkg(args)
	if err != nil {
		return nil, err
	}
	// Write to file.
	outDir, err := detectOutputDir(pkg.GoFiles)
	if err != nil {
		return nil, err
	}

	// Inspect package
	serviceFile, serviceDecl, serviceSpec, serviceType, rpcMethods := Inspect(pkg, serviceName)
	if serviceFile == nil || serviceDecl == nil || serviceSpec == nil || serviceType == nil {
		log.Fatal("error: not found service")
	}

	serviceInfo, err := ParseServiceInfo(serviceDecl)
	if err != nil {
		return nil, err
	}
	serviceInfo.SetServiceName(serviceName)
	serviceInfo.SetOutDir(outDir)
	serviceInfo.SetPkgPath(pkg.PkgPath)
	serviceInfo.SetPackageName(pkg.Name)

	imports := ExtractGoImports(serviceFile)
	serviceInfo.SetImports(imports)
	routers, err := ParseRouterInfos(rpcMethods, serviceInfo, pathToLower)
	if err != nil {
		log.Fatal(err)
	}
	serviceInfo.SetRouters(routers)
	return serviceInfo, nil
}

type ServiceInfo struct {
	Name        string
	Description string
	BasePath    string
	Routers     []*RouterInfo
	FullName    string
	PackageName string
	OutDir      string
	PkgPath     string
	Imports     map[string]*GoImport
}

func (info *ServiceInfo) SetServiceName(s string) {
	info.Name = s
}

func (info *ServiceInfo) SetRouters(routers []*RouterInfo) {
	info.Routers = routers
}

func (info *ServiceInfo) SetFullName(name string) {
	info.FullName = name
}

func (info *ServiceInfo) SetPackageName(name string) {
	info.PackageName = name
	info.FullName = fmt.Sprintf("%s.%s", info.PackageName, info.Name)
}

func (info *ServiceInfo) SetOutDir(dir string) {
	info.OutDir = dir
}

func (info *ServiceInfo) SetPkgPath(pkgPath string) {
	info.PkgPath = pkgPath
}

func (info *ServiceInfo) SetImports(imports map[string]*GoImport) {
	info.Imports = imports
}

var ErrPathInvalid = errors.New("path invalid")

func NewService(comments []string) (*ServiceInfo, error) {
	var basePath string
	desc := &bytes.Buffer{}
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		if seg[0] != GORS {
			_, _ = fmt.Fprint(desc, text, " ")
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(Path)):
				v, ok := ExtractValue(s, Path)
				if !ok {
					return nil, ErrPathInvalid
				}
				basePath = path.Join(basePath, v)
			case strings.HasPrefix(s, GORS):
				continue
			case "" == s:
				continue
			}
		}
	}
	return &ServiceInfo{Description: desc.String(), BasePath: basePath}, nil
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
