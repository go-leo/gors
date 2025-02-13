package internal

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

type Service struct {
	ProtoService *protogen.Service
	Endpoints    []*Endpoint
}

func (s *Service) Unexported(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func (s *Service) FullName() string {
	return string(s.ProtoService.Desc.FullName())
}

func (s *Service) Name() string {
	return s.ProtoService.GoName
}

func (s *Service) AppendGorillaRouteName() string {
	return "Append" + s.Name() + "GorillaRoute"
}

func (s *Service) GorillaServiceName() string {
	return s.Name() + "GorillaService"
}

func (s *Service) ServerEndpointsName() string {
	return s.Name() + "ServerEndpoints"
}

func (s *Service) TransportsName() string {
	return s.Name() + "Transports"
}

func (s *Service) ServerName() string {
	return s.Name() + "Server"
}

func (s *Service) GorillaName() string {
	return s.Name() + "Gorilla"
}

func (s *Service) GorillaTransportsName() string {
	return s.GorillaName() + "Transports"
}

func (s *Service) HttpRoutesName() string {
	return s.Name() + "HttpRoutes"
}

func (s *Service) GorillaRoutesName() string {
	return s.Name() + "GorillaRoutes"
}

func (s *Service) UnimplementedServerName() string {
	return "Unimplemented" + s.ProtoService.GoName + "Server"
}

func (s *Service) GorillaRequestDecoderName() string {
	return s.GorillaName() + "RequestDecoder"
}

func (s *Service) GorillaResponseEncoderName() string {
	return s.GorillaName() + "ResponseEncoder"
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			ProtoService: pbService,
		}
		var endpoints []*Endpoint
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("leo: unsupport stream method, %s", endpoint.FullName())
			}
			endpoint.SetHttpRule()
			endpoints = append(endpoints, endpoint)
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}
