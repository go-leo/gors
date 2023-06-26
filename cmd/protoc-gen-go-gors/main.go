package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var pathToLower *bool
var restful *bool

func main() {
	var flags flag.FlagSet
	pathToLower = flags.Bool("path_to_lower", false, "make path to lower case")
	restful = flags.Bool("restful", false, "use restful service")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
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
