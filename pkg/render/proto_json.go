package render

import (
	"encoding/json"
	"github.com/gin-gonic/gin/render"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

var _ render.Render = ProtoJSON{}

var jsonContentType = []string{"application/json; charset=utf-8"}

type ProtoJSON struct {
	Data           any
	MarshalOptions protojson.MarshalOptions
}

func (r ProtoJSON) Render(w http.ResponseWriter) (err error) {
	if httpBody, ok := r.Data.(*httpbody.HttpBody); ok {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{httpBody.GetContentType()}
		}
		_, err := w.Write(httpBody.GetData())
		return err
	}
	if response, ok := r.Data.(*rpchttp.HttpResponse); ok {
		w.WriteHeader(int(response.GetStatus()))
		for _, header := range response.GetHeaders() {
			w.Header().Set(header.GetKey(), header.GetValue())
		}
		_, err := w.Write(response.GetBody())
		return err
	}
	r.WriteContentType(w)
	m, ok := r.Data.(proto.Message)
	if !ok {
		jsonBytes, err := json.Marshal(r.Data)
		if err != nil {
			return err
		}
		_, err = w.Write(jsonBytes)
		return err
	}
	jsonBytes, err := r.MarshalOptions.Marshal(m)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (r ProtoJSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
}
