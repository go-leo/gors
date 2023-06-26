package demo

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

//go:generate gors -service CustomBinderRender

// CustomBinderRender
// @GORS @Path(/api/CustomBinderRender)
type CustomBinderRender interface {
	// Custom
	// @GORS @POST @Path(/Custom) @CustomBinding @CustomRender
	Custom(ctx context.Context, req *CustomReq) (*CustomResp, error)
}

type CustomReq struct {
	V string
}

func (req *CustomReq) Bind(c *gin.Context) error {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	req.V = string(data)
	return nil
}

type CustomResp struct {
}

func (resp *CustomResp) Render(c *gin.Context) {
	data, _ := json.Marshal(resp)
	c.Data(http.StatusOK, "text/plain", data)
}
