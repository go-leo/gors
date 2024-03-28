package parser

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"strings"
)

type ServiceInfo struct {
	Name        string
	Description string
	BasePath    string
	Routers     []*RouterInfo
	FullName    string
	PackageName string
}

func (info *ServiceInfo) SetServiceName(s string) {
	info.Name = s
}

func (info *ServiceInfo) SetRouters(routers []*RouterInfo) {
	info.Routers = routers
}

func (info *ServiceInfo) SetFullName(name string) {
	info.FullName = name
}

func (info *ServiceInfo) SetPackageName(name string) {
	info.PackageName = name
	info.FullName = fmt.Sprintf("%s.%s", info.PackageName, info.Name)
}

var ErrPathInvalid = errors.New("path invalid")

func NewService(comments []string) (*ServiceInfo, error) {
	var basePath string
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
					return nil, ErrPathInvalid
				}
				basePath = path.Join(basePath, v)
			case strings.HasPrefix(s, GORS):
				continue
			case "" == s:
				continue
			}
		}
	}
	return &ServiceInfo{Description: desc.String(), BasePath: basePath}, nil
}
