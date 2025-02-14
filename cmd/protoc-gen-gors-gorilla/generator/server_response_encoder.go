package generator

import (
	"fmt"
	"github.com/go-leo/gors/v2/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

func (f *Generator) GenerateServerResponseEncoder(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.GorillaResponseEncoderName(), " struct {")
	g.P("marshalOptions ", internal.ProtoJsonPackage.Ident("MarshalOptions"))
	g.P("unmarshalOptions ", internal.ProtoJsonPackage.Ident("UnmarshalOptions"))
	g.P("}")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		g.P("func (encoder ", service.GorillaResponseEncoderName(), ")", endpoint.Name(), "(ctx ", internal.ContextIdent, ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", resp *", endpoint.OutputGoIdent(), ") error {")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			srcValue := []any{"resp"}
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintGoogleApiHttpBodyEncodeBlock(g, srcValue)
			default:
				f.PrintJsonEncodeBlock(g, srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.Get", bodyField.GoName, "()"}
			if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
				f.PrintGoogleApiHttpBodyEncodeBlock(g, srcValue)
			} else {
				f.PrintJsonEncodeBlock(g, srcValue)
			}
		}
		g.P("return nil")
		g.P("}")
	}
	g.P()
	return nil
}

func (f *Generator) PrintGoogleApiHttpBodyEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"w.Header().Set(", strconv.Quote("Content-Type"), ", "}, srcValue...), ".GetContentType())")...)
	g.P(append(append([]any{"for _, src := range "}, srcValue...), ".GetExtensions() {")...)
	g.P("dst, err := ", internal.AnypbPackage.Ident("UnmarshalNew"), "(src, encoder.unmarshalOptions)")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P("metadata, ok := dst.(*", internal.StructpbPackage.Ident("Struct"), ")")
	g.P("if !ok {")
	g.P("continue")
	g.P("}")
	g.P("for key, value := range metadata.GetFields() {")
	g.P("w.Header().Add(key, string(", internal.ErrorxPackage.Ident("Ignore"), "(encoder.marshalOptions(value))))")
	g.P("}")
	g.P("}")
	g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	g.P(append(append([]any{"if ", "_, err := w.Write("}, srcValue...), ".GetData())", "; err != nil {")...)
	g.P("return err")
	g.P("}")
}

func (f *Generator) PrintJsonEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", strconv.Quote(internal.JsonContentType), ")")
	g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	g.P(append(append([]any{"data, err := ", "encoder.marshalOptions.Marshal("}, srcValue...), ")")...)
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P("if _, err := w.Write(data); err != nil {")
	g.P("return err")
	g.P("}")
}
