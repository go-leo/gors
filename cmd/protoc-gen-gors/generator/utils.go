package generator

import (
	"fmt"
	openapiv3 "github.com/google/gnostic/openapiv3"
	"google.golang.org/protobuf/compiler/protogen"
	"net/http"
	"regexp"
	"strings"
)

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

var pathPattern = regexp.MustCompile("{([^=}]+)}")

func convertPath(namedPath string) string {
	if allMatches := pathPattern.FindAllStringSubmatch(namedPath, -1); allMatches != nil {
		for _, matches := range allMatches {
			for i := 0; i < len(matches); i += 2 {
				namedPath = strings.Replace(namedPath, matches[i], ":"+matches[i+1], 1)
			}
		}
	}
	return namedPath
}

func convertHttpMethod(pathItem *openapiv3.PathItem) (string, *openapiv3.Operation) {
	switch {
	case pathItem.GetGet() != nil:
		return http.MethodGet, pathItem.GetGet()
	case pathItem.GetPost() != nil:
		return http.MethodPost, pathItem.GetPost()
	case pathItem.GetPut() != nil:
		return http.MethodPut, pathItem.GetPut()
	case pathItem.GetDelete() != nil:
		return http.MethodDelete, pathItem.GetDelete()
	case pathItem.GetPatch() != nil:
		return http.MethodPatch, pathItem.GetPatch()
	}
	return "", nil
}
