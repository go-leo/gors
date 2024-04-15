package gors

import "google.golang.org/protobuf/reflect/protoreflect"

// Names for google.protobuf.Any.
const (
	anyMessageName     protoreflect.Name     = "Any"
	anyMessageFullName protoreflect.FullName = "google.protobuf.Any"
)

// isMessageSet returns whether the message uses the MessageSet wire format.
func isMessageSet(md protoreflect.MessageDescriptor) bool {
	xmd, ok := md.(interface{ IsMessageSet() bool })
	return ok && xmd.IsMessageSet()
}
