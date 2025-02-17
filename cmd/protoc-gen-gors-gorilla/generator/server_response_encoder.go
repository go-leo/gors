package generator

import (
	"fmt"
	"github.com/go-leo/gors/v2/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (f *Generator) GenerateServerResponseEncoder(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.GorillaResponseEncoderName()), " struct {")
	g.P("marshalOptions ", internal.ProtoJsonMarshalOptionsIdent)
	g.P("unmarshalOptions ", internal.ProtoJsonUnmarshalOptionsIdent)
	g.P("}")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		g.P("func (encoder ", service.Unexported(service.GorillaResponseEncoderName()), ")", endpoint.Name(), "(ctx ", internal.ContextIdent, ", w ", internal.ResponseWriterIdent, ", resp *", endpoint.OutputGoIdent(), ") error {")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			srcValue := []any{"resp"}
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintHttpBodyEncodeBlock(g, srcValue)
			case "google.rpc.HttpResponse":
				f.PrintHttpResponseEncodeBlock(g, srcValue)
			default:
				f.PrintResponseEncodeBlock(g, srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.Get", bodyField.GoName, "()"}
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.PrintHttpBodyEncodeBlock(g, srcValue)
				default:
					f.PrintResponseEncodeBlock(g, srcValue)
				}
			}
		}
		g.P("}")
	}
	g.P()
	return nil
}

func (f *Generator) PrintHttpBodyEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", internal.HttpBodyEncoderIdent, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *Generator) PrintHttpResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", internal.HttpResponseEncoderIdent, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *Generator) PrintResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", internal.ResponseEncoderIdent, "(ctx, w, "}, srcValue...), ", encoder.marshalOptions)")...)
}
