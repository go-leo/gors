package generator

import (
	"github.com/go-leo/gors/v2/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

func (f *Generator) GenerateServerRequestDecoder(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.RequestDecoderName()), " struct {")
	g.P("unmarshalOptions ", gen.ProtoJsonUnmarshalOptionsIdent)
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (decoder ", service.Unexported(service.RequestDecoderName()), ")", endpoint.Name(), "(ctx ", gen.ContextIdent, ", r *", gen.RequestIdent, ") (*", endpoint.InputGoIdent(), ", error){")
		g.P("req := &", endpoint.InputGoIdent(), "{}")
		bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
		if err != nil {
			return err
		}

		if bodyMessage != nil {
			switch bodyMessage.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintHttpBodyDecodeBlock(g, []any{"req"})
			case "google.rpc.HttpRequest":
				f.PrintHttpRequestEncodeBlock(g, []any{"req"})
			default:
				f.PrintRequestDecodeBlock(g, []any{"req"})
			}
		} else if bodyField != nil {
			tgtValue := []any{"req.", bodyField.GoName}
			g.P(append(append(append([]any{"if "}, tgtValue...), " == nil {"))...)
			g.P(append(tgtValue, " = &", bodyField.Message.GoIdent, "{}")...)
			g.P("}")
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.PrintHttpBodyDecodeBlock(g, tgtValue)
				default:
					f.PrintRequestDecodeBlock(g, []any{"req.", bodyField.GoName})
				}
			}
		}

		if len(namedPathFields)+len(pathFields) > 0 {
			g.P("vars := ", gen.UrlxPackage.Ident("FormFromMap"), "(", gen.VarsIdent, "(r)", ")")
			f.PrintNamedPathField(g, namedPathFields, endpoint.HttpRule())
			f.PrintPathField(g, pathFields)
		}

		if len(queryFields) > 0 {
			g.P("queries := r.URL.Query()")
			g.P("var queryErr error")
			f.PrintQueryField(g, queryFields)
			g.P("if queryErr != nil {")
			g.P("return nil, queryErr")
			g.P("}")
		}

		g.P("return req, nil")
		g.P("}")
	}
	g.P()
	return nil
}

func (f *Generator) PrintHttpBodyDecodeBlock(g *protogen.GeneratedFile, tgtValue []any) {
	g.P(append(append([]any{"if err := ", gen.HttpBodyDecoderIdent, "(ctx, r, "}, tgtValue...), "); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}

func (f *Generator) PrintHttpRequestEncodeBlock(g *protogen.GeneratedFile, tgtValue []any) {
	g.P(append(append([]any{"if err := ", gen.HttpRequestDecoderIdent, "(ctx, r, "}, tgtValue...), "); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}

func (f *Generator) PrintRequestDecodeBlock(g *protogen.GeneratedFile, tgtValue []any) {
	g.P(append(append([]any{"if err := ", gen.RequestDecoderIdent, "(ctx, r, "}, tgtValue...), ", decoder.unmarshalOptions); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}

func (f *Generator) PrintNamedPathField(g *protogen.GeneratedFile, namedPathFields []*protogen.Field, httpRule *gen.HttpRule) {
	for i, namedPathField := range namedPathFields {
		fullFieldName := gen.FullFieldName(namedPathFields[:i+1])
		if i < len(namedPathFields)-1 {
			g.P("if req.", fullFieldName, " == nil {")
			g.P("req.", fullFieldName, " = &", namedPathField.Message.GoIdent, "{}")
			g.P("}")
		} else {
			_, _, namedPathTemplate, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
			tgtValue := []any{"req.", fullFieldName, " = "}
			srcValue := []any{gen.SprintfIdent, "(", strconv.Quote(namedPathTemplate)}
			for _, namedPathParameter := range namedPathParameters {
				srcValue = append(srcValue, ", vars.Get(", strconv.Quote(namedPathParameter), ")")
			}
			srcValue = append(srcValue, ")")

			switch namedPathField.Desc.Kind() {
			case protoreflect.StringKind:
				f.PrintStringValueAssign(g, tgtValue, srcValue, namedPathField.Desc.HasPresence())
			case protoreflect.MessageKind:
				f.PrintWrapStringValueAssign(g, tgtValue, srcValue)
			}
		}
	}
}

func (f *Generator) PrintPathField(g *protogen.GeneratedFile, pathFields []*protogen.Field) {
	if len(pathFields) <= 0 {
		return
	}
	form := "vars"
	errName := "varErr"
	g.P("var ", errName, " error")
	for _, field := range pathFields {
		fieldName := string(field.Desc.Name())

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", ", errName, " = "}
		srcValue := []any{"vars.Get(", strconv.Quote(fieldName), ")"}

		goType, pointer := gen.FieldGoType(g, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt"), fieldName, form, errName)
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUintPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint"), fieldName, form, errName)
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt"), fieldName, form, errName)
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUintPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint"), fieldName, form, errName)
			}
		case protoreflect.FloatKind: // float32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloatPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat"), fieldName, form, errName)
			}
		case protoreflect.DoubleKind: // float64
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloatPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat"), fieldName, form, errName)
			}
		case protoreflect.StringKind: // string
			f.PrintStringValueAssign(g, tgtValue, srcValue, pointer)
		case protoreflect.EnumKind: // enum int32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntPtr["+g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt["+g.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
			case "google.protobuf.FloatValue":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
			case "google.protobuf.Int64Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
			case "google.protobuf.UInt64Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
			case "google.protobuf.Int32Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
			case "google.protobuf.UInt32Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
			case "google.protobuf.BoolValue":
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
			case "google.protobuf.StringValue":
				f.PrintWrapStringValueAssign(g, tgtValue, srcValue)
			}
		}
	}
	g.P("if ", errName, " != nil {")
	g.P("return nil, ", errName)
	g.P("}")
}

func (f *Generator) PrintQueryField(g *protogen.GeneratedFile, queryFields []*protogen.Field) {
	for _, field := range queryFields {
		fieldName := string(field.Desc.Name())

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", queryErr = "}
		srcValue := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
		if field.Desc.IsList() {
			srcValue = []any{"queries[", strconv.Quote(fieldName), "]"}
		}

		goType, pointer := gen.FieldGoType(g, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		form := "queries"
		errName := "queryErr"

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBoolSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
				}
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUintSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUintPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint"), fieldName, form, errName)
				}
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUintSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUintPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint"), fieldName, form, errName)
				}
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloatSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloatPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat"), fieldName, form, errName)
				}
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloatSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloatPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat"), fieldName, form, errName)
				}
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintStringListAssign(g, tgtValue, srcValue)
			} else {
				f.PrintStringValueAssign(g, tgtValue, srcValue, pointer)
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntSlice["+g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetIntPtr["+g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt["+g.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
				}
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBoolValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, gen.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintWrapStringListAssign(g, tgtValue, srcValue)
				} else {
					f.PrintWrapStringValueAssign(g, tgtValue, srcValue)
				}
			}
		}
	}

}

func (f *Generator) PrintFieldAssign(g *protogen.GeneratedFile, tgtValue []any, goType []any, getter protogen.GoIdent, key string, form string, errName string) {
	g.P(append(append([]any{}, tgtValue...), append(append([]any{gen.FormDecoderIdent, "["}, goType...), append([]any{"](", errName, ", ", form, ", ", strconv.Quote(key), ", ", getter}, ")")...)...)...)
}

func (f *Generator) PrintStringValueAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any, hasPresence bool) {
	if hasPresence {
		g.P(append(tgtValue, append(append([]any{gen.ProtoStringIdent, "("}, srcValue...), ")")...)...)
	} else {
		g.P(append(tgtValue, srcValue...)...)
	}
}

func (f *Generator) PrintWrapStringValueAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(tgtValue, append(append([]any{gen.WrapperspbStringIdent, "("}, srcValue...), ")")...)...)
}

func (f *Generator) PrintStringListAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(tgtValue, srcValue...)...)
}

func (f *Generator) PrintWrapStringListAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(tgtValue, append(append([]any{gen.WrapStringSliceIdent, "("}, srcValue...), ")")...)...)
}
