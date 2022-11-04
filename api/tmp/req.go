package tmp

type SumRequest struct {
	A int64 `form:"a" json:"a" xml:"a" uri:"a" header:"a" binding:"required"`
	B int64 `form:"b" json:"b" xml:"b" uri:"b" header:"b" binding:"required"`
}

type SumReply struct {
	V int64
}
