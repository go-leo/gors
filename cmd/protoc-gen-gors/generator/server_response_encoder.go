package generator

import (
	"fmt"
	"github.com/go-leo/gors/v2/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (f *Generator) GenerateServerResponseEncoder(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.ResponseEncoderName()), " struct {")
	g.P("marshalOptions ", gen.ProtoJsonMarshalOptionsIdent)
	g.P("unmarshalOptions ", gen.ProtoJsonUnmarshalOptionsIdent)
	g.P("responseTransformer ", gen.ResponseTransformerIdent)
	g.P("}")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		g.P("func (encoder ", service.Unexported(service.ResponseEncoderName()), ")", endpoint.Name(), "(ctx ", gen.ContextIdent, ", w ", gen.ResponseWriterIdent, ", resp *", endpoint.OutputGoIdent(), ") error {")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				srcValue := []any{"resp"}
				f.PrintHttpBodyEncodeBlock(g, srcValue)
			case "google.rpc.HttpResponse":
				srcValue := []any{"resp"}
				f.PrintHttpResponseEncodeBlock(g, srcValue)
			default:
				srcValue := []any{"encoder.responseTransformer(ctx, resp)"}
				f.PrintResponseEncodeBlock(g, srcValue)
			}
		default:
			bodyField := gen.FindField(bodyParameter, endpoint.Output())
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
	g.P(append(append([]any{"return ", gen.HttpBodyEncoderIdent, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *Generator) PrintHttpResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", gen.HttpResponseEncoderIdent, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *Generator) PrintResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", gen.ResponseEncoderIdent, "(ctx, w, "}, srcValue...), ", encoder.marshalOptions)")...)
}
