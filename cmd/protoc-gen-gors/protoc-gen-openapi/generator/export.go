package generator

import (
	v3 "github.com/google/gnostic/openapiv3"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"strings"
)

var GlobalOpenAPIv3Generator = NewOpenAPIv3Generator(nil, Configuration{}, nil)

func (g *OpenAPIv3Generator) FullMessageTypeName(message protoreflect.MessageDescriptor) string {
	return g.reflect.fullMessageTypeName(message)
}

func (g *OpenAPIv3Generator) BuildDocumentV3() *v3.Document {
	return g.buildDocumentV3()
}

func (g *OpenAPIv3Generator) newDocument(file *protogen.File) *v3.Document {
	d := &v3.Document{}

	d.Openapi = "3.0.3"
	d.Info = &v3.Info{
		Version:     *g.conf.Version,
		Title:       *g.conf.Title,
		Description: *g.conf.Description,
	}

	d.Paths = &v3.Paths{}
	d.Components = &v3.Components{
		Schemas: &v3.SchemasOrReferences{
			AdditionalProperties: []*v3.NamedSchemaOrReference{},
		},
	}

	// Merge any `Document` annotations with the current
	extDocument := proto.GetExtension(file.Desc.Options(), v3.E_Document)
	if extDocument != nil {
		proto.Merge(d, extDocument.(*v3.Document))
	}
	return d
}

func (g *OpenAPIv3Generator) ParseHttpRule(rule *annotations.HttpRule) (string, string, string) {
	var methodName string
	var path string
	var body string

	body = rule.Body
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		methodName = "GET"
	case *annotations.HttpRule_Post:
		path = pattern.Post
		methodName = "POST"
	case *annotations.HttpRule_Put:
		path = pattern.Put
		methodName = "PUT"
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		methodName = "DELETE"
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		methodName = "PATCH"
	case *annotations.HttpRule_Custom:
		path = "custom-unsupported"
	default:
		path = "unknown-unsupported"
	}
	return methodName, path, body
}

func (g *OpenAPIv3Generator) BuildPaths(file *protogen.File, service *protogen.Service, method *protogen.Method) ([]string, *v3.Paths, []string) {
	d := g.newDocument(file)

	comment := g.filterCommentString(method.Comments.Leading)
	inputMessage := method.Input
	outputMessage := method.Output
	operationID := service.GoName + "_" + method.GoName

	rules := make([]*annotations.HttpRule, 0)

	extHTTP := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
	if extHTTP != nil && extHTTP != annotations.E_Http.InterfaceOf(annotations.E_Http.Zero()) {
		rule := extHTTP.(*annotations.HttpRule)
		rules = append(rules, rule)
		rules = append(rules, rule.AdditionalBindings...)
	}

	var paths []string
	var bodies []string
	for _, rule := range rules {
		var path string
		var methodName string
		var body string

		body = rule.Body
		switch pattern := rule.Pattern.(type) {
		case *annotations.HttpRule_Get:
			path = pattern.Get
			methodName = "GET"
		case *annotations.HttpRule_Post:
			path = pattern.Post
			methodName = "POST"
		case *annotations.HttpRule_Put:
			path = pattern.Put
			methodName = "PUT"
		case *annotations.HttpRule_Delete:
			path = pattern.Delete
			methodName = "DELETE"
		case *annotations.HttpRule_Patch:
			path = pattern.Patch
			methodName = "PATCH"
		case *annotations.HttpRule_Custom:
			path = "custom-unsupported"
		default:
			path = "unknown-unsupported"
		}

		if methodName != "" {
			defaultHost := proto.GetExtension(service.Desc.Options(), annotations.E_DefaultHost).(string)

			op, path2 := g.buildOperationV3(
				d, operationID, service.GoName, comment, defaultHost, path, body, inputMessage, outputMessage)

			// Merge any `Operation` annotations with the current
			extOperation := proto.GetExtension(method.Desc.Options(), v3.E_Operation)
			if extOperation != nil {
				proto.Merge(op, extOperation.(*v3.Operation))
			}

			g.addOperationToDocumentV3(d, op, path2, methodName)
			paths = append(paths, path)
			bodies = append(bodies, body)
		}
	}

	return paths, d.Paths, bodies
}

func (g *OpenAPIv3Generator) AddPathsToDocumentV3(file *protogen.File, service *protogen.Service) {
	document := g.newDocument(file)
	g.addPathsToDocumentV3(document, []*protogen.Service{service})
	return
}

func (g *OpenAPIv3Generator) FindAndFormatFieldName(name string, inMessage *protogen.Message) string {
	return g.findAndFormatFieldName(name, inMessage)
}

type PathParameters struct {
	ParameterOrReference  *v3.ParameterOrReference
	ParameterOrReferences []*v3.ParameterOrReference
	// Name book.name
	Name string
	// Parameters shelf,book
	Parameters []string
	// Template shelves/%s/books/%s
	Template string
	// json or proto
	Naming string
}

// FindSimplePathParameters Find simple path parameters like {id}
func (g *OpenAPIv3Generator) FindSimplePathParameters(path string, inputMessage *protogen.Message) (string, []*PathParameters) {
	// Initialize the list of operation parameters.
	parameters := make([]*PathParameters, 0)
	if allMatches := g.pathPattern.FindAllStringSubmatch(path, -1); allMatches != nil {
		for _, matches := range allMatches {
			// Add the value to the list of covered parameters.
			pathParameter := g.findAndFormatFieldName(matches[1], inputMessage)
			path = strings.Replace(path, matches[1], pathParameter, 1)

			// Add the path parameters to the operation parameters.
			var fieldSchema *v3.SchemaOrReference

			var fieldDescription string
			field := g.findField(pathParameter, inputMessage)
			if field != nil {
				fieldSchema = g.reflect.schemaOrReferenceForField(field.Desc)
				fieldDescription = g.filterCommentString(field.Comments.Leading)
			} else {
				// If field does not exist, it is safe to set it to string, as it is ignored downstream
				fieldSchema = &v3.SchemaOrReference{
					Oneof: &v3.SchemaOrReference_Schema{
						Schema: &v3.Schema{
							Type: "string",
						},
					},
				}
			}

			parameters = append(parameters, &PathParameters{
				ParameterOrReference: &v3.ParameterOrReference{
					Oneof: &v3.ParameterOrReference_Parameter{
						Parameter: &v3.Parameter{
							Name:        pathParameter,
							In:          "path",
							Description: fieldDescription,
							Required:    true,
							Schema:      fieldSchema,
						},
					},
				},
				Naming: *g.conf.Naming,
			})
		}
	}
	return path, parameters
}

// FindNamedPathParameters Find named path parameters like {name=shelves/*}
func (g *OpenAPIv3Generator) FindNamedPathParameters(path string, inputMessage *protogen.Message) (string, *PathParameters) {
	// Initialize the list of operation parameters.
	var parameters *PathParameters

	if matches := g.namedPathPattern.FindStringSubmatch(path); matches != nil {
		parameters = &PathParameters{
			Naming: *g.conf.Naming,
		}
		// Build a list of named path parameters.
		namedPathParameters := make([]string, 0)

		// Convert the path from the starred form to use named path parameters.
		starredPath := matches[2]
		parts := strings.Split(starredPath, "/")
		// The starred path is assumed to be in the form "things/*/otherthings/*".
		// We want to convert it to "things/{thingsId}/otherthings/{otherthingsId}".
		for i := 0; i < len(parts)-1; i += 2 {
			section := parts[i]
			namedPathParameter := g.findAndFormatFieldName(section, inputMessage)
			namedPathParameter = singular(namedPathParameter)
			parts[i+1] = "{" + namedPathParameter + "}"
			namedPathParameters = append(namedPathParameters, namedPathParameter)
		}
		// Rewrite the path to use the path parameters.
		newPath := strings.Join(parts, "/")

		path = strings.Replace(path, matches[0], newPath, 1)

		parameters.Name = matches[1]
		parameters.Parameters = namedPathParameters
		parameters.Template = strings.Replace(starredPath, "*", "%s", -1)
		// Add the named path parameters to the operation parameters.
		for _, namedPathParameter := range namedPathParameters {
			parameters.ParameterOrReferences = append(parameters.ParameterOrReferences,
				&v3.ParameterOrReference{
					Oneof: &v3.ParameterOrReference_Parameter{
						Parameter: &v3.Parameter{
							Name:        namedPathParameter,
							In:          "path",
							Required:    true,
							Description: "The " + namedPathParameter + " id.",
							Schema: &v3.SchemaOrReference{
								Oneof: &v3.SchemaOrReference_Schema{
									Schema: &v3.Schema{
										Type: "string",
									},
								},
							},
						},
					},
				})
		}
	}

	return path, parameters
}

func (g *OpenAPIv3Generator) FindBodyParameters(bodyField string, inputMessage *protogen.Message) (*v3.RequestBodyOrReference, error) {
	if bodyField == "" {
		return nil, nil
	}
	// If a body field is specified, we need to pass a message as the request body.
	var requestSchema *v3.SchemaOrReference

	if bodyField == "*" {
		// Pass the entire request message as the request body.
		requestSchema = g.reflect.schemaOrReferenceForMessage(inputMessage.Desc)

	} else {
		// If body refers to a message field, use that type.
		for _, field := range inputMessage.Fields {
			if string(field.Desc.Name()) == bodyField {
				switch field.Desc.Kind() {
				case protoreflect.StringKind:
					requestSchema = &v3.SchemaOrReference{
						Oneof: &v3.SchemaOrReference_Schema{
							Schema: &v3.Schema{
								Type: "string",
							},
						},
					}

				case protoreflect.MessageKind:
					requestSchema = g.reflect.schemaOrReferenceForMessage(field.Message.Desc)

				default:
					log.Printf("unsupported field type %+v", field.Desc)
				}
				break
			}
		}
	}

	return &v3.RequestBodyOrReference{
		Oneof: &v3.RequestBodyOrReference_RequestBody{
			RequestBody: &v3.RequestBody{
				Required: true,
				Content: &v3.MediaTypes{
					AdditionalProperties: []*v3.NamedMediaType{{Name: "application/json", Value: &v3.MediaType{Schema: requestSchema}}},
				},
			}}}, nil
}
