package parser

import (
	"log"
	"path"
	"strings"
)

type ServiceInfo struct {
	Name     string
	BasePath string
	Routers  []*RouterInfo
}

func NewService(name string, comments []string) *ServiceInfo {
	info := &ServiceInfo{Name: name}
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		if seg[0] != GORS {
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(Path)):
				v, ok := ExtractValue(s, Path)
				if !ok {
					log.Fatalf("error: %s path invalid", s)
				}
				info.BasePath = path.Join(info.BasePath, v)
			case strings.HasPrefix(s, GORS):
				continue
			case "" == s:
				continue
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	return info
}
