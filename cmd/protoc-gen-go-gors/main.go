package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "v1.4.0"

var pathToLower *bool
var requireUnimplemented *bool

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-gors %v\n", version)
		return
	}

	var flags flag.FlagSet
	pathToLower = flags.Bool("path_to_lower", false, "make path to lower case")
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")
	protogen.Options{ParamFunc: flags.Set}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}
