package main

import (
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
	return
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

	serviceInfo, err := parser.ParseService(args, *serviceName, *pathToLower)
	if err != nil {
		log.Fatal(err)
	}

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
	outputPath := filepath.Join(serviceInfo.OutDir, fmt.Sprintf("%s_swagger.yaml", strings.ToLower(*serviceName)))
	if err := os.WriteFile(outputPath, swaggerYaml, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}

	log.Printf("%s.%s wrote %s", serviceInfo.PackageName, *serviceName, outputPath)
}
