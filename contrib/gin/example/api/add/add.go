package add

import "context"

type AddService interface {
	// Sum
	// #HTTP @Method("GET","POST")
	// #HTTP @Path("/")
	// #HTTP @Consume("json")
	// #HTTP @Produce("")
	Sum(context.Context, *SumRequest) (*SumReply, error)
	// Concat
	// #HTTP @Method("GET","POST")
	// #HTTP @Path("/")
	// #HTTP @Consumes("","")
	// #HTTP @Produces("","")
	Concat(context.Context, *ConcatRequest) (*ConcatReply, error)
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
