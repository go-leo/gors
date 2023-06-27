package binding

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

var ProtoJSON = protoJSONBinding{}

type protoJSONBinding struct{}

func (b protoJSONBinding) Name() string {
	return "json"
}

func (b protoJSONBinding) Bind(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return b.BindBody(body, obj)
}

func (b protoJSONBinding) BindBody(body []byte, obj any) error {
	m, ok := obj.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to convert data, %v", obj)
	}
	return protojson.Unmarshal(body, m)
}
