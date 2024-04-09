package main

import (
	"flag"
	"github.com/go-leo/gors/cmd/protoc-gen-go-gorsopenapi/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"path/filepath"
	"strings"
)

var flags flag.FlagSet

func main() {
	conf := generator.Configuration{
		Version:         flags.String("version", "0.0.1", "version number text, e.g. 1.2.3"),
		Title:           flags.String("title", "", "name of the API"),
		Description:     flags.String("description", "", "description of the API"),
		Naming:          flags.String("naming", "json", `naming convention. Use "proto" for passing names directly from the proto files`),
		FQSchemaNaming:  flags.Bool("fq_schema_naming", false, `schema naming convention. If "true", generates fully-qualified schema names by prefixing them with the proto message package name`),
		EnumType:        flags.String("enum_type", "integer", `type for enum serialization. Use "string" for string-based serialization`),
		CircularDepth:   flags.Int("depth", 2, "depth of recursion for circular messages"),
		DefaultResponse: flags.Bool("default_response", true, `add default response. If "true", automatically adds a default response to operations which use the google.rpc.Status message. Useful if you use envoy or grpc-gateway to transcode as they use this type for their default error responses.`),
		PathToLower:     flags.Bool("path_to_lower", false, "make path to lower case"),
	}
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			for _, service := range file.Services {
				outfileName := strings.TrimSuffix(file.Desc.Path(), filepath.Ext(file.Desc.Path())) + "_" + service.GoName + ".openapi.yaml"
				outputFile := plugin.NewGeneratedFile(outfileName, "")
				gen := generator.NewOpenAPIv3Generator(plugin, conf, file, service)
				if err := gen.Run(outputFile); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
