package main

import (
	"flag"
	"fmt"
	"github.com/go-leo/gnostic/cmd/protoc-gen-jsonschema/generator"
	"github.com/go-leo/gors/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"strings"
)

var pathToLower *bool
var requireUnimplemented *bool

func main() {

	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-gors %v\n", internal.Version)
		return
	}

	builder := strings.Builder{}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	builder.WriteString(fmt.Sprintln("wd"))
	builder.WriteString(fmt.Sprintln(dir))
	builder.WriteString(fmt.Sprintln())
	builder.WriteString(fmt.Sprintln("Environ"))
	environ := os.Environ()
	for _, s := range environ {
		builder.WriteString(fmt.Sprintln(s))
	}
	builder.WriteString(fmt.Sprintln())
	builder.WriteString(fmt.Sprintln("Args"))
	for _, arg := range os.Args {
		builder.WriteString(fmt.Sprintln(arg))
	}

	err = os.WriteFile("./sdd.txt", []byte(builder.String()), 0666)
	if err != nil {
		panic(err)
	}

	var flags flag.FlagSet
	pathToLower = flags.Bool("path_to_lower", false, "make path to lower case")
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	conf := generator.Configuration{
		BaseURL:  flags.String("baseurl", "", "the base url to use in schema ids"),
		Version:  flags.String("version", "http://json-schema.org/draft-07/schema#", "schema version URL used in $schema. Currently supported: draft-06, draft-07"),
		Naming:   flags.String("naming", "json", `naming convention. Use "proto" for passing names directly from the proto files`),
		EnumType: flags.String("enum_type", "integer", `type for enum serialization. Use "string" for string-based serialization`),
	}

	opts := protogen.Options{ParamFunc: flags.Set}

	opts.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		g := generator.NewJSONSchemaGenerator(gen, conf)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			schemas := g.BuildSchemasFromMessages(f.Messages)
			for _, schema := range schemas {
				//fmt.Println(schema.Name)
				_ = schema
			}
		}
		return nil
	})
}
