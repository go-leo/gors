package parser

import (
	"bytes"
	"fmt"
	"github.com/go-openapi/spec"
	"log"
	"path"
	"strings"
)

type ServiceInfo struct {
	Name        string
	Description string
	BasePath    string
	Routers     []*RouterInfo
}

func (info *ServiceInfo) Swagger() *spec.Swagger {
	schemes := []string{"http", "https"}

	var consumes []string
	var produces []string
	paths := &spec.Paths{
		VendorExtensible: spec.VendorExtensible{},
		Paths:            map[string]spec.PathItem{},
	}
	definitions := spec.Definitions{}
	parameters := map[string]spec.Parameter{}
	responses := map[string]spec.Response{}
	for _, router := range info.Routers {
		consumes = append(consumes, router.BindingContentType)
		produces = append(produces, router.RenderContentType)

		var parameters []spec.Parameter
		for _, binding := range router.Bindings {
			switch binding {
			case ReaderBinding:
			case BytesBinding:
			case StringBinding:
			case UriBinding:
				for _, param := range router.PathParams() {
					parameters = append(parameters, *spec.PathParam(param))
				}
			case QueryBinding:
				for _, param := range router.QueryParams() {
					parameters = append(parameters, *spec.QueryParam(param))
				}
			case HeaderBinding:
				for _, param := range router.HeaderParams() {
					parameters = append(parameters, *spec.FormDataParam(param))
				}
			case JSONBinding:
			case XMLBinding:
			case FormBinding:
				for _, param := range router.FormParams() {
					parameters = append(parameters, *spec.FormDataParam(param))
				}
			case FormPostBinding:
				for _, param := range router.FormParams() {
					parameters = append(parameters, *spec.FormDataParam(param))
				}
			case FormMultipartBinding:
				for _, param := range router.FormParams() {
					parameters = append(parameters, *spec.FormDataParam(param))
				}
				for _, param := range router.FileParams() {
					parameters = append(parameters, *spec.FileParam(param))
				}
			case ProtoBufBinding:
			case MsgPackBinding:
			case YAMLBinding:
			case TOMLBinding:
			case CustomBinding:
			case ProtoJSONBinding:

			}
		}

		pathItemProps := spec.PathItemProps{}
		operation := &spec.Operation{
			VendorExtensible: spec.VendorExtensible{},
			OperationProps: spec.OperationProps{
				Description:  router.Description,
				Consumes:     []string{router.BindingContentType},
				Produces:     []string{router.RenderContentType},
				Schemes:      schemes,
				Tags:         []string{},
				Summary:      "",
				ExternalDocs: &spec.ExternalDocumentation{},
				ID:           "",
				Deprecated:   false,
				Security:     []map[string][]string{},
				Parameters:   parameters,
				Responses: &spec.Responses{
					VendorExtensible: spec.VendorExtensible{},
					ResponsesProps: spec.ResponsesProps{
						Default:             nil,
						StatusCodeResponses: map[int]spec.Response{},
					},
				},
			},
		}
		switch router.HttpMethod {
		case GET:
			pathItemProps.Get = operation
		case POST:
			pathItemProps.Post = operation
		case PUT:
			pathItemProps.Put = operation
		case DELETE:
			pathItemProps.Delete = operation
		case PATCH:
			pathItemProps.Patch = operation
		case HEAD:
			pathItemProps.Head = operation
		case CONNECT:
		case OPTIONS:
		case TRACE:
		default:
		}
		paths.Paths[router.Path] = spec.PathItem{
			Refable:          spec.Refable{},
			VendorExtensible: spec.VendorExtensible{},
			PathItemProps:    pathItemProps,
		}
	}
	metadata := &spec.Info{
		VendorExtensible: spec.VendorExtensible{
			Extensions: spec.Extensions{},
		},
		InfoProps: spec.InfoProps{
			Description:    info.Description,
			Title:          info.Name,
			TermsOfService: "",
			Contact:        nil,
			License:        nil,
			Version:        "",
		},
	}

	return &spec.Swagger{
		VendorExtensible: spec.VendorExtensible{
			Extensions: map[string]interface{}{},
		},
		SwaggerProps: spec.SwaggerProps{
			ID:                  info.Name,
			Consumes:            consumes,
			Produces:            produces,
			Schemes:             schemes,
			Swagger:             "2.0",
			Info:                metadata,
			Host:                "",
			BasePath:            info.BasePath,
			Paths:               paths,
			Definitions:         definitions,
			Parameters:          parameters,
			Responses:           responses,
			SecurityDefinitions: spec.SecurityDefinitions{},
			Security:            []map[string][]string{},
			Tags:                []spec.Tag{},
			ExternalDocs:        &spec.ExternalDocumentation{},
		},
	}
}

func (info *ServiceInfo) SetServiceName(s string) {
	info.Name = s
}

func (info *ServiceInfo) SetRouters(routers []*RouterInfo) {
	info.Routers = routers
}

func NewService(comments []string) *ServiceInfo {
	info := &ServiceInfo{}
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
					log.Fatalf("error: %s path invalid", s)
				}
				info.BasePath = path.Join(info.BasePath, v)
			case strings.HasPrefix(s, GORS):
				continue
			case "" == s:
				continue
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	info.Description = desc.String()
	return info
}
