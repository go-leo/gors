package generator

import (
	"errors"
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// singular produces the singular form of a collection name.
func singular(plural string) string {
	if strings.HasSuffix(plural, "ves") {
		return strings.TrimSuffix(plural, "ves") + "f"
	}
	if strings.HasSuffix(plural, "ies") {
		return strings.TrimSuffix(plural, "ies") + "y"
	}
	if strings.HasSuffix(plural, "s") {
		return strings.TrimSuffix(plural, "s")
	}
	return plural
}

func getValueKind(message protoreflect.MessageDescriptor) string {
	valueField := getValueField(message)
	return valueField.Kind().String()
}

func getValueField(message protoreflect.MessageDescriptor) protoreflect.FieldDescriptor {
	fields := message.Fields()
	return fields.ByName("value")
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func getHttpRules(method *protogen.Method) []*annotations.HttpRule {
	rules := make([]*annotations.HttpRule, 0)
	extHTTP := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
	if extHTTP != nil && extHTTP != annotations.E_Http.InterfaceOf(annotations.E_Http.Zero()) {
		rule := extHTTP.(*annotations.HttpRule)
		rules = append(rules, rule)
		rules = append(rules, rule.AdditionalBindings...)
	}
	return rules
}

var httpMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

func getHttpMethod(rule *annotations.HttpRule) (string, string, error) {
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return pattern.Get, http.MethodGet, nil
	case *annotations.HttpRule_Post:
		return pattern.Post, http.MethodPost, nil
	case *annotations.HttpRule_Put:
		return pattern.Put, http.MethodPut, nil
	case *annotations.HttpRule_Delete:
		return pattern.Delete, http.MethodDelete, nil
	case *annotations.HttpRule_Patch:
		return pattern.Patch, http.MethodPatch, nil
	case *annotations.HttpRule_Custom:
		return pattern.Custom.Path, pattern.Custom.Kind, nil
	default:
		return "", "", errors.New("unknown-unsupported")
	}
}

func findField(name string, inMessage *protogen.Message) *protogen.Field {
	for _, field := range inMessage.Fields {
		if string(field.Desc.Name()) == name || string(field.Desc.JSONName()) == name {
			return field
		}
	}
	return nil
}

func findAndFormatFieldName(name string, inMessage *protogen.Message) string {
	field := findField(name, inMessage)
	if field != nil {
		return formatFieldName(field.Desc)
	}

	return name
}
