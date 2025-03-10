package main

import (
	"flag"
	"fmt"
	"github.com/go-leo/gors/cmd/protoc-gen-gors/generator"
	openapigenerator "github.com/go-leo/gors/cmd/protoc-gen-gors/protoc-gen-openapi/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var flags flag.FlagSet

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-gors %v\n", "1.5.8")
		return
	}

	conf := openapigenerator.Configuration{
		DocVersion:      flags.String("doc_version", "0.0.1", "version number text, e.g. 1.2.3"),
		Title:           flags.String("title", "", "name of the API"),
		Description:     flags.String("description", "", "description of the API"),
		Naming:          flags.String("naming", "json", `naming convention. Use "proto" for passing names directly from the proto files`),
		FQSchemaNaming:  flags.Bool("fq_schema_naming", false, `schema naming convention. If "true", generates fully-qualified schema names by prefixing them with the proto message package name`),
		EnumType:        flags.String("enum_type", "integer", `type for enum serialization. Use "string" for string-based serialization`),
		CircularDepth:   flags.Int("depth", 2, "depth of recursion for circular messages"),
		DefaultResponse: flags.Bool("default_response", true, `add default response. If "true", automatically adds a default response to operations which use the google.rpc.Status message. Useful if you use envoy or grpc-gateway to transcode as they use this type for their default error responses.`),
		OutputMode:      flags.String("output_mode", "merged", `output generation mode. By default, a single openapi.yaml is generated at the out folder. Use "source_relative' to generate a separate '[inputfile].openapi.yaml' next to each '[inputfile].proto'.`),
	}

	generator.GrpcServer = flags.Bool("grpc_server", false, `add grpc server.`)
	generator.GrpcClient = flags.Bool("grpc_client", false, `add grpc client.`)

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			filename := file.GeneratedFilenamePrefix + "_gors.pb.go"
			outputFile := plugin.NewGeneratedFile(filename, file.GoImportPath)
			openapiGenerator := openapigenerator.NewOpenAPIv3Generator(plugin, conf, []*protogen.File{file})
			gen := generator.NewGenerator(plugin, file, outputFile, openapiGenerator)
			if err := gen.Run(); err != nil {
				return err
			}
		}
		return nil
	})
}
