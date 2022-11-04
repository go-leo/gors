package add

import (
	"context"
)

//go:generate gors -service Add

// Add
// @GORS @Path("/api/add")  @Path("/api") @Consume("JSON") @Produce("JSON")
type Add interface {
	// Concat
	// @GORS @Method("GET")
	// @GORS @POST
	// @GORS @Path("/bytes")
	// @GORS @Consume("Form")
	// @GORS @Produce("JSON")
	Bytes(context.Context, []byte) (string, error)
	/*
		// Sum
		// @GORS @Method("GET")
		// @GORS @HEAD
		// @GORS @Path("/Sum")
		// @GORS @Consume("JSON")
		// @GORS @Produce("JSON")
		Sum(context.Context, *SumRequest) (*SumReply, error)
		// Concat
		// @GORS @Method("GET")
		// @GORS @POST
		// @GORS @Path("/Concat")
		// @GORS @Consume("Form")
		// @GORS @Produce("JSON")
		Concat(context.Context, *ConcatRequest) (*ConcatReply, error)

		// Concat
		// @GORS @Method("GET")
		// @GORS @POST
		// @GORS @Path("/string")
		// @GORS @Consume("Form")
		// @GORS @Produce("JSON")
		String(context.Context, string) (string, error)

		// Concat
		// @GORS @Method("GET")
		// @GORS @POST
		// @GORS @Path("/otherpkg")
		// @GORS @Consume("Form")
		// @GORS @Produce("JSON")
		OtherPkg(context.Context, *tmp.SumRequest) (*tmp.SumReply, error)
		// Concat
		// @GORS @Method("GET")
		// @GORS @POST
		// @GORS @Path("/otherpkg2")
		// @GORS @Consume("Form")
		// @GORS @Produce("JSON")
		OtherPkg2(context.Context, tmp.SumRequest) (tmp.SumReply, error)

		// Concat
		// @GORS @Method("GET")
		// @GORS @POST
		// @GORS @Path("/otherpkg2")
		// @GORS @Consume("Form")
		// @GORS @Produce("JSON")
		OtherPkg3(tmp.SumReply, tmp.SumRequest) (tmp.SumReply, error)*/
}

type SumRequest struct {
	A int64 `form:"a" json:"a" xml:"a" uri:"a" header:"a" binding:"required"`
	B int64 `form:"b" json:"b" xml:"b" uri:"b" header:"b" binding:"required"`
}

type SumReply struct {
	V int64
}

type ConcatRequest struct {
	A string
	B string
}

type ConcatReply struct {
	V string
}
