package generator

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func serviceName(service *protogen.Service) string {
	return service.GoName + "Service"
}

func gRPCServerName(service *protogen.Service) string {
	return service.GoName + "Server"
}

func gRPCClientName(service *protogen.Service) string {
	return service.GoName + "Client"
}

func serviceWrapperName(service *protogen.Service) string {
	return "_" + serviceName(service) + "Wrapper"
}

func gRPCServerWrapperName(service *protogen.Service) string {
	return "_" + gRPCServerName(service) + "Wrapper"
}

func grpcClientWrapperName(service *protogen.Service) string {
	return "_" + gRPCClientName(service) + "Wrapper"
}

func fullMethodName(service *protogen.Service, method *protogen.Method) string {
	return fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())
}

func handlerName(service *protogen.Service, method *protogen.Method) string {
	return fmt.Sprintf("_%s_%s_GORS_Handler", service.GoName, method.GoName) //"_ProtoDemo_POSTProtoJSONBindingProtoJSONRender_GORS_Handler"
}

func serviceRoutesFunctionName(service *protogen.Service) string {
	return serviceName(service) + "Routes"
}

func gRPCClientRoutesFunctionName(service *protogen.Service) string {
	return gRPCClientName(service) + "Routes"
}

func gRPCServerRoutesFunctionName(service *protogen.Service) string {
	return gRPCServerName(service) + "Routes"
}
