package main

import (
	"context"
	"github.com/go-leo/gox/netx/httpx"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrailers(t *testing.T) {
	receiver, err := outgoing.Sender().Post().
		URLString("http://localhost:8088/v1/POSTSetHeaderTrailer").
		Header("TE", "Trailers").
		JSONBody(&protodemo.HelloRequest{
			Name:   "Jax",
			Age:    10,
			Salary: 2.3,
			Token:  "xxx",
		}).Send(context.Background(), httpx.PooledClient())
	assert.NoError(t, err)
	t.Log(receiver.Headers())
	t.Log(receiver.TextBody())
	t.Log(receiver.Trailers())
}
