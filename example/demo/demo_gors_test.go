package demo

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"maps"
	"net/http"
	"slices"
	"time"
)

func Example() {
	router := mux.NewRouter()
	router = AppendDemoGorsRoute(router, NewMockDemo())
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

type MockDemo struct {
	m map[int64]string
}

func (svc *MockDemo) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	fmt.Println("CreateUser:", request.GetName())
	id := time.Now().Unix()
	svc.m[id] = request.GetName()
	return &CreateUserResponse{
		Item: &User{
			Id:   id,
			Name: request.GetName(),
		}}, nil
}

func (svc *MockDemo) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*emptypb.Empty, error) {
	fmt.Println("DeleteUser:", request.GetId())
	delete(svc.m, request.GetId())
	return &emptypb.Empty{}, nil
}

func (svc *MockDemo) ModifyUser(ctx context.Context, request *ModifyUserRequest) (*emptypb.Empty, error) {
	fmt.Println("ModifyUser:", request.GetId(), request.GetName())
	svc.m[request.GetId()] = request.GetName()
	return &emptypb.Empty{}, nil
}

func (svc *MockDemo) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	fmt.Println("GetUser:", request.GetId())
	return &GetUserResponse{Item: &User{
		Id:   request.GetId(),
		Name: svc.m[request.GetId()],
	}}, nil
}

func (svc *MockDemo) ListUser(ctx context.Context, request *ListUserRequest) (*ListUserResponse, error) {
	fmt.Println("ListUser:", request.GetPageNum(), request.GetPageSize())
	keys := maps.Keys(svc.m)
	ids := slices.SortedFunc(keys, func(a int64, b int64) int {
		return int(a - b)
	})
	resp := &ListUserResponse{List: make([]*User, 0)}
	for i := request.GetPageSize() * (request.GetPageNum() - 1); i < int64(len(ids)) && int64(len(resp.List)) < request.GetPageSize(); i++ {
		id := ids[i]
		resp.List = append(resp.List, &User{
			Id:   id,
			Name: svc.m[id],
		})
	}
	return resp, nil
}

func NewMockDemo() DemoGorsService {
	return &MockDemo{m: map[int64]string{}}
}
