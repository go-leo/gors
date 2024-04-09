package generator

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// contains returns true if an array contains a specified string.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// appendUnique appends a string, to a string slice, if the string is not already in the slice
func appendUnique(s []string, e string) []string {
	if !contains(s, e) {
		return append(s, e)
	}
	return s
}

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
