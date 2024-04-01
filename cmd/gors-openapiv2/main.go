package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/go-leo/gors/cmd/internal"
	"github.com/go-leo/gors/internal/pkg/parser"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
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
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Usage = Usage
	flag.Parse()
	if *showVersion {
		fmt.Printf("gors %v\n", internal.Version)
		return
	}

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
	pkg, err := parser.LoadPkg(args)
	if err != nil {
		log.Fatal(err)
	}
	pkgName := pkg.Name
	goFiles := pkg.GoFiles
	pkgPath := pkg.PkgPath

	// Inspect package
	serviceFile, serviceDecl, serviceSpec, serviceType, rpcMethods := parser.Inspect(pkg, *serviceName)
	if serviceFile == nil || serviceDecl == nil || serviceSpec == nil || serviceType == nil {
		log.Fatal("error: not found service")
	}

	serviceInfo, err := parser.ParseServiceInfo(serviceDecl)
	if err != nil {
		log.Fatal(err)
	}
	serviceInfo.SetServiceName(*serviceName)
	serviceInfo.SetPackageName(pkgName)
	imports := parser.ExtractGoImports(serviceFile)
	routers, err := parser.ParseRouterInfos(rpcMethods, imports, serviceInfo, *pathToLower)
	if err != nil {
		log.Fatal(err)
	}
	serviceInfo.SetRouters(routers)

	swagger, err := serviceInfo.Swagger()
	if err != nil {
		log.Fatal(err)
	}
	swaggerJson, err := swagger.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	var swaggerObj interface{}
	if err := yaml.Unmarshal(swaggerJson, &swaggerObj); err != nil {
		log.Fatal(err)
	}
	swaggerYaml, err := yaml.Marshal(swaggerObj)
	if err != nil {
		log.Fatal(err)
	}

	// Write to file.
	outDir, err := detectOutputDir(goFiles)
	if err != nil {
		log.Fatalf("error: detect output dir: %s", err)
	}
	outputPath := filepath.Join(outDir, fmt.Sprintf("%s_swagger.yaml", strings.ToLower(*serviceName)))
	if err := os.WriteFile(outputPath, swaggerYaml, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}

	log.Printf("%s.%s wrote %s", pkgPath, *serviceName, outputPath)
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
