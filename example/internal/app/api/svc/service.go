package svc

import (
	"context"
	stderr "errors"
	"fmt"

	pkgerr "github.com/pkg/errors"

	"github.com/go-leo/gors/example/api/demo"
	"github.com/go-leo/gors/example/api/errors"
)

var _ demo.Service = new(Service)

type Service struct{}

func (svc *Service) Method(ctx context.Context, req *demo.MethodReq) (*demo.MethodResp, error) {
	fmt.Println(req.ID)
	if req.ID == 0 {
		e := errors.ErrUnknown
		e.Cause = stderr.New("cause error")
		fmt.Printf("%+v\n", e)
		return nil, e
	}
	if req.ID == 1 {
		e := errors.ErrUnknown
		e.Cause = pkgerr.WithStack(stderr.New("cause error"))
		fmt.Printf("%+v\n", e)
		return nil, e
	}
	if req.ID == 2 {
		e := testStackErr()
		fmt.Printf("%+v\n", e)
		return nil, e
	}
	if req.ID == 3 {
		// e := testStackErrWithCode()
		e := errors.ErrUnknown
		fmt.Printf("%+v\n", e)
		return nil, e
	}
	return &demo.MethodResp{V: 10}, nil
}

// func testCauseErr() error {
// 	dberr := stderr.New("db error")
// 	return errors.NewErrUserNotFound().WithCause(dberr)
// }

func testStackErr() error { // 还是要实现
	dberr := stderr.New("db error")
	return pkgerr.WithStack(dberr)
}

// func testStackErrWithCode() error { // 还是要实现
// 	dberr := stderr.New("db error")
// 	dberr = pkgerr.WithStack(dberr)
// 	return errors.NewErrUserNotFound("not found uid:%d", 1).WithCause(dberr)
// }
