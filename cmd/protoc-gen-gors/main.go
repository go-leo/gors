package main

import (
	"flag"
	"github.com/go-leo/gors/cmd/protoc-gen-gors/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var flags flag.FlagSet

func main() {
	generator.Naming = flags.String("naming", "json", `naming convention. Use "proto" for passing names directly from the proto files`)
	generator.FQSchemaNaming = flags.Bool("fq_schema_naming", false, `schema naming convention. If "true", generates fully-qualified schema names by prefixing them with the proto message package name`)
	generator.EnumType = flags.String("enum_type", "integer", `type for enum serialization. Use "string" for string-based serialization`)
	generator.CircularDepth = flags.Int("depth", 2, "depth of recursion for circular messages")
	generator.DefaultResponse = flags.Bool("default_response", true, `add default response. If "true", automatically adds a default response to operations which use the google.rpc.Status message. Useful if you use envoy or grpc-gateway to transcode as they use this type for their default error responses.`)

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		// Enable "optional" keyword in front of type (e.g. optional string label = 1;)
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			filename := file.GeneratedFilenamePrefix + "_gors.go"
			outputFile := plugin.NewGeneratedFile(filename, file.GoImportPath)
			gen := generator.NewOpenAPIv3Generator(plugin, file, outputFile)
			if err := gen.Run(); err != nil {
				return err
			}
		}
		return nil
	})
}
