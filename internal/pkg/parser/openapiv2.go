package parser

import (
	"github.com/go-openapi/spec"
)

const swaggerVersion = "2.0"

func (info *ServiceInfo) Swagger() (*spec.Swagger, error) {
	host := ""
	pathsDoc, err := info.PathsDoc()
	if err != nil {
		return nil, err
	}
	return &spec.Swagger{
		VendorExtensible: spec.VendorExtensible{
			Extensions: map[string]interface{}{},
		},
		SwaggerProps: spec.SwaggerProps{
			ID:                  info.Name,
			Consumes:            info.consumesDoc(),
			Produces:            info.producesDoc(),
			Schemes:             info.schemesDoc(),
			Swagger:             swaggerVersion,
			Info:                info.InfoDoc(),
			Host:                host,
			BasePath:            info.BasePath,
			Paths:               pathsDoc,
			Definitions:         info.definitionsDoc(),
			Parameters:          info.parameterDoc(),
			Responses:           info.responseDoc(),
			SecurityDefinitions: spec.SecurityDefinitions{},
			Security:            []map[string][]string{},
			Tags:                []spec.Tag{},
			ExternalDocs:        &spec.ExternalDocumentation{},
		},
	}, nil
}

func (info *ServiceInfo) consumesDoc() []string {
	var consumes []string
	for _, router := range info.Routers {
		contentType := router.BindingContentType
		if contentType == "" {
			continue
		}
		consumes = append(consumes, contentType)
	}
	return consumes
}

func (info *ServiceInfo) producesDoc() []string {
	var produces []string
	for _, router := range info.Routers {
		contentType := router.RenderContentType
		if contentType == "" {
			continue
		}
		produces = append(produces, contentType)
	}
	return produces
}

func (info *ServiceInfo) responseDoc() map[string]spec.Response {
	return map[string]spec.Response{}
}

func (info *ServiceInfo) parameterDoc() map[string]spec.Parameter {
	return map[string]spec.Parameter{}
}

func (info *ServiceInfo) definitionsDoc() spec.Definitions {
	return spec.Definitions{}
}

func (info *ServiceInfo) InfoDoc() *spec.Info {
	return &spec.Info{
		VendorExtensible: spec.VendorExtensible{Extensions: spec.Extensions{}},
		InfoProps: spec.InfoProps{
			Description:    info.Description,
			Title:          info.Name,
			TermsOfService: "",
			Contact:        nil,
			License:        nil,
			Version:        "",
		},
	}
}

func (info *ServiceInfo) schemesDoc() []string {
	schemes := []string{"http", "https"}
	return schemes
}

func (info *ServiceInfo) PathsDoc() (*spec.Paths, error) {
	paths := &spec.Paths{
		VendorExtensible: spec.VendorExtensible{
			Extensions: map[string]interface{}{},
		},
		Paths: map[string]spec.PathItem{},
	}
	for _, router := range info.Routers {
		pathDoc, err := router.PathDoc()
		if err != nil {
			return nil, err
		}
		paths.Paths[router.Path] = pathDoc
	}
	return paths, nil
}

func (router *RouterInfo) PathDoc() (spec.PathItem, error) {
	pathItemProps, err := router.PathItemProps()
	if err != nil {
		return spec.PathItem{}, err
	}
	return spec.PathItem{
		PathItemProps: pathItemProps,
	}, nil
}

func (router *RouterInfo) PathItemProps() (spec.PathItemProps, error) {
	props := spec.PathItemProps{}
	doc, err := router.OperationDoc(GET)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Get = doc

	doc, err = router.OperationDoc(PUT)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Put = doc

	doc, err = router.OperationDoc(POST)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Post = doc

	doc, err = router.OperationDoc(DELETE)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Delete = doc

	doc, err = router.OperationDoc(OPTIONS)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Options = doc

	doc, err = router.OperationDoc(HEAD)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Head = doc

	doc, err = router.OperationDoc(PATCH)
	if err != nil {
		return spec.PathItemProps{}, err
	}
	props.Patch = doc
	return props, nil
}
