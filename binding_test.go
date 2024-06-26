package gors_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/tests"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"net/http"
	"testing"
	"time"
)

func TestPayloadBinding(t *testing.T) {
	relativePath := "/demo/:message_id/shelves/:shelf/books/:book"
	urlString := "http://localhost:8080/demo/10/shelves/20/books/30?page_size=10&page_token=xxl"

	engine := gin.New()
	engine.POST(relativePath, func(context *gin.Context) {
		ctx := gors.NewContext(context, "demo")
		binding := gors.ProtoHttpBinding(&gors.ProtoHttpParameter{
			Path: []*gors.PathParameter{
				{Name: "message_id", Type: "string"},
			},
			NamedPath: &gors.NamedPathParameter{
				Name:       "book.name",
				Parameters: []string{"shelf", "book"},
				Template:   "shelves/%s/books/%s",
			},
			Query: []*gors.QueryParameter{
				{Name: "page_size", Type: "integer"},
				{Name: "page_token", Type: "string"},
			},
			Body: &gors.BodyParameter{
				Name: "sub_message",
				Type: "object",
			},
		})
		var req tests.Message
		err := gors.RequestBind(ctx, &req, "", binding)
		if err != nil {
			panic(err)
		}
	})

	go func() {
		if err := engine.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	t.Log("send...")
	subMessage := tests.SubMessage{
		Address: "shanghai",
		SubSubMessage: &tests.SubSubMessage{
			Title: "leo",
			Text:  "leo, gors",
		},
	}
	send, err := outgoing.Sender().
		Post().
		URLString(urlString).
		JSONBody(&subMessage).
		Send(context.Background(), &http.Client{})
	if err != nil {
		panic(err)
	}
	body, err := send.TextBody()
	if err != nil {
		panic(err)
	}
	t.Log(body)

	select {}
}
