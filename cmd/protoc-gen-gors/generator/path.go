package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
	"strings"
)

var (
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
	namedPathPattern = regexp.MustCompile("{(.+)=(.+)}")
)

// Find simple path parameters like {id}
func findSimplePathParameters(path string, inputMessage *protogen.Message) (string, []string, []string) {
	var pathParameters []string
	var coveredParameters []string
	if allMatches := pathPattern.FindAllStringSubmatch(path, -1); allMatches != nil {
		for _, matches := range allMatches {
			// Add the value to the list of covered parameters.
			coveredParameters = append(coveredParameters, matches[1])
			pathParameter := findAndFormatFieldName(matches[1], inputMessage)
			pathParameters = append(pathParameters, pathParameter)
			path = strings.Replace(path, "{"+matches[1]+"}", ":"+pathParameter, 1)
		}
	}
	return path, pathParameters, coveredParameters
}

// Find named path parameters like {name=shelves/*}
func findNamedPathParameters(path string, inputMessage *protogen.Message) (string, []string, []string) {
	var pathParameters []string
	var coveredParameters []string

	if matches := namedPathPattern.FindStringSubmatch(path); matches != nil {
		// Build a list of named path parameters.
		namedPathParameters := make([]string, 0)

		// Add the "name=" "name" value to the list of covered parameters.
		coveredParameters = append(coveredParameters, matches[1])
		// Convert the path from the starred form to use named path parameters.
		starredPath := matches[2]
		parts := strings.Split(starredPath, "/")
		// The starred path is assumed to be in the form "things/*/otherthings/*".
		// We want to convert it to "things/{thingsId}/otherthings/{otherthingsId}".
		for i := 0; i < len(parts)-1; i += 2 {
			section := parts[i]
			namedPathParameter := findAndFormatFieldName(section, inputMessage)
			namedPathParameter = singular(namedPathParameter)
			pathParameters = append(pathParameters, namedPathParameter)
			parts[i+1] = ":" + namedPathParameter
			namedPathParameters = append(namedPathParameters, namedPathParameter)
		}
		// Rewrite the path to use the path parameters.
		newPath := strings.Join(parts, "/")
		path = strings.Replace(path, matches[0], newPath, 1)
	}
	return path, pathParameters, coveredParameters
}

func findPathParameters(path string, bodyField string, inputMessage *protogen.Message) (string, []string) {
	// coveredParameters tracks the parameters that have been used in the body or path.
	coveredParameters := make([]string, 0)
	if bodyField != "" {
		coveredParameters = append(coveredParameters, bodyField)
	}
	pathParameters := make([]string, 0)
	path, pps, cpps := findSimplePathParameters(path, inputMessage)
	coveredParameters = append(coveredParameters, cpps...)
	pathParameters = append(pathParameters, pps...)

	path, pps, cpps = findNamedPathParameters(path, inputMessage)
	coveredParameters = append(coveredParameters, cpps...)
	pathParameters = append(pathParameters, pps...)

	return path, pathParameters
}
