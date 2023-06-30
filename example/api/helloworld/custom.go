package helloworld

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-leo/gors"
	"github.com/go-leo/gox/convx"
	"io"
	"net/http"
)

func (req *HelloRequest) Bind(ctx context.Context) error {
	c := gors.FromContext(ctx)
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	segs := bytes.Split(data, []byte(" "))
	for _, seg := range segs {
		subSeg := bytes.Split(seg, []byte(":"))
		if len(subSeg) != 2 {
			return fmt.Errorf("%s invalid", string(seg))
		}
		switch string(subSeg[0]) {
		case "name":
			req.Name = string(subSeg[1])
		case "age":
			req.Age = convx.ToInt32(subSeg[1])
		case "salary":
			req.Salary = convx.ToFloat64(subSeg[1])
		case "Token":
			req.Token = string(subSeg[1])
		}
	}
	return nil
}

func (reply *HelloReply) Render(ctx context.Context) {
	c := gors.FromContext(ctx)
	c.String(http.StatusOK, "message:%s", reply.GetMessage())
}
