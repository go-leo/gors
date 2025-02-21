# gors 
gors 参考 Google API 设计方案，通过定义 proto 文件，快速生成Go语言的Restful服务路由。
gors 底层基于 [gorilla](http://github.com/gorilla/mux) 管理http路由。

# 路由规则
详细定义规则见 googleapi 文档 [http.proto](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)

# 安装
```
go install github.com/go-leo/gors/v2/cmd/protoc-gen-gors@latest
```

# Example
```protobuf
syntax = "proto3";
package leo.gors.demo.v1;
option go_package = "github.com/go-leo/gors/v2/example/demo/v1;demo";

import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service Demo {

  // CreateUser 创建用户
  // `POST /v1/user { "name": "Leo" }` | `CreateUserRequest(name: "Leo")`
  rpc CreateUser (CreateUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post : "/v1/user"
      body : "*"
    };
  }

  // DeleteUser 删除用户
  // `DELETE /v1/user/10000 | `DeleteUserRequest(id: 10000)`
  rpc DeleteUser (DeleteUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete : "/v1/user/{id}"
    };
  }

  // ModifyUser 修改用户
  // `PUT /v1/user/10000 { "name": "Leo" }` | `ModifyUserRequest(id: 10000, name: "Leo")`
  rpc ModifyUser (ModifyUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put : "/v1/user/{id}"
      body : "*"
    };
  }

  // GetUser 获取用户
  // `GET /v1/user/10000` | `GetUserRequest(id: 10000)`
  rpc GetUser (GetUserRequest) returns (GetUserResponse) {
    option (google.api.http) = {
      get : "/v1/user/{id}"
    };
  }

  // ListUser 获取用户列表
  // `GET /v1/users?page_num=1&page_size=10` | `ListUserRequest(page_num: 1, page_size: 10)`
  rpc ListUser (ListUserRequest) returns (ListUserResponse) {
    option (google.api.http) = {
      get : "/v1/users"
    };
  }
}

message User {
  int64 id = 1;
  string name = 2;
}

message CreateUserRequest {
  string name = 1;
}

message DeleteUserRequest {
  int64 id = 1;
}

message ModifyUserRequest {
  int64 id = 1;
  string name = 2;
}

message GetUserRequest {
  int64 id = 1;
}

message GetUserResponse {
  User item = 1;
}

message ListUserRequest {
  int64 page_num = 1;
  int64 page_size = 2;
}

message ListUserResponse {
  repeated User list = 1;
}
```

# 生成路由
```
protoc \
--proto_path=. \
--proto_path=../third_party \
--proto_path=../../ \
--go_out=. \
--go_opt=paths=source_relative \
--gors-gorilla_out=. \
--gors-gorilla_opt=paths=source_relative \
demo.proto
```

```
demo
├── demo.pb.go
├── demo.proto
└── demo_gors.gorilla.pb.go
```

# 代码
```go
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
```

更多示例见 [example]https://github.com/go-leo/gors/tree/v2/example

