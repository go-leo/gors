package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/protobuftypes"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/mathx/randx"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net"
	"net/http"
	"time"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, protobuftypes.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8097")
	if err != nil {
		panic(err)
	}
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}

type MessagingService struct {
}

func (m MessagingService) CreateMessage(ctx context.Context, message *protobuftypes.Message) (*protobuftypes.Message, error) {
	return message, nil
}

func (m MessagingService) CreateMessagesFromCSV(ctx context.Context, body *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return body, nil
}

func (m MessagingService) ListMessages(ctx context.Context, empty *emptypb.Empty) (*structpb.Value, error) {
	i := randx.Int()
	switch i % 6 {
	case 0:
		return structpb.NewNullValue(), nil
	case 1:
		return structpb.NewBoolValue(true), nil
	case 2:
		return structpb.NewNumberValue(randx.Float64()), nil
	case 3:
		return structpb.NewStringValue(randx.WordString(12)), nil
	case 4:
		newStruct, _ := structpb.NewStruct(map[string]interface{}{
			"name": randx.WordString(12),
			"id":   randx.Int(),
		})
		return structpb.NewStructValue(newStruct), nil
	case 5:
		list, _ := structpb.NewList([]any{1, 2, 3, 4, 5})
		return structpb.NewListValue(list), nil
	default:
		return nil, errors.New(convx.ToString(i))
	}
}

func (m MessagingService) ListMessagesCSV(ctx context.Context, body *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return body, nil
}

func (m MessagingService) GetMessage(ctx context.Context, message *protobuftypes.Message) (*protobuftypes.Message, error) {
	newStruct, _ := structpb.NewStruct(map[string]interface{}{
		"name": randx.WordString(12),
		"id":   randx.Int(),
	})
	_, _ = structpb.NewList([]any{1, 2, 3, 4, 5})
	return &protobuftypes.Message{
		MessageId:  message.GetMessageId(),
		StringType: randx.WordString(10),
		RecursiveType: &protobuftypes.RecursiveParent{
			ParentId: randx.Int31(),
			Child: &protobuftypes.RecursiveChild{
				ChildId: randx.Int31(),
				Parent: &protobuftypes.RecursiveParent{
					ParentId: randx.Int31(),
					Child: &protobuftypes.RecursiveChild{
						ChildId: randx.Int31(),
					},
				},
			},
		},
		EmbeddedType: &protobuftypes.Message_EmbMessage{MessageId: randx.WordString(32)},
		SubType: &protobuftypes.SubMessage{
			MessageId: randx.HexString(12),
			SubSubMessage: &protobuftypes.SubSubMessage{
				MessageId: randx.HexString(16),
				Integers:  []int32{randx.Int31(), randx.Int31(), randx.Int31()},
			},
		},
		RepeatedType: []string{randx.WordString(12), randx.WordString(12), randx.WordString(12)},
		RepeatedSubType: []*protobuftypes.SubMessage{
			{
				MessageId: randx.HexString(12),
				SubSubMessage: &protobuftypes.SubSubMessage{
					MessageId: randx.HexString(16),
					Integers:  []int32{randx.Int31(), randx.Int31(), randx.Int31()},
				},
			},
			{
				MessageId: randx.HexString(12),
				SubSubMessage: &protobuftypes.SubSubMessage{
					MessageId: randx.HexString(16),
					Integers:  []int32{randx.Int31(), randx.Int31(), randx.Int31()},
				},
			},
		},
		RepeatedRecursiveType: []*protobuftypes.RecursiveParent{
			{
				ParentId: randx.Int31(),
				Child: &protobuftypes.RecursiveChild{
					ChildId: randx.Int31(),
					Parent: &protobuftypes.RecursiveParent{
						ParentId: randx.Int31(),
						Child: &protobuftypes.RecursiveChild{
							ChildId: randx.Int31(),
						},
					},
				},
			},
			{
				ParentId: randx.Int31(),
				Child: &protobuftypes.RecursiveChild{
					ChildId: randx.Int31(),
					Parent: &protobuftypes.RecursiveParent{
						ParentId: randx.Int31(),
						Child: &protobuftypes.RecursiveChild{
							ChildId: randx.Int31(),
						},
					},
				},
			},
		},
		MapType: map[string]string{
			randx.WordString(12): randx.WordString(12),
			randx.WordString(12): randx.WordString(12),
			randx.WordString(12): randx.WordString(12),
			randx.WordString(12): randx.WordString(12),
			randx.WordString(12): randx.WordString(12),
			randx.WordString(12): randx.WordString(12),
		},
		Body:      newStruct,
		Media:     []*structpb.Struct{newStruct, newStruct, newStruct},
		NotUsed:   new(emptypb.Empty),
		ValueType: errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
		RepeatedValueType: []*structpb.Value{
			errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
			errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
			errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
			errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
			errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
			errorx.Ignore(m.ListMessages(ctx, &emptypb.Empty{})),
		},
		BoolValueType:   wrapperspb.Bool(randx.Bool()),
		BytesValueType:  wrapperspb.Bytes([]byte(randx.WordString(33))),
		Int32ValueType:  wrapperspb.Int32(randx.Int31()),
		Uint32ValueType: wrapperspb.UInt32(randx.Uint32()),
		StringValueType: wrapperspb.String(randx.HexString(32)),
		Int64ValueType:  wrapperspb.Int64(randx.Int63()),
		Uint64ValueType: wrapperspb.UInt64(randx.Uint64()),
		FloatValueType:  wrapperspb.Float(randx.Float32()),
		DoubleValueType: wrapperspb.Double(randx.Float64()),
		TimestampType:   timestamppb.Now(),
		DurationType:    durationpb.New(time.Hour),
	}, nil
}

func (m MessagingService) UpdateMessage(ctx context.Context, message *protobuftypes.Message) (*structpb.Struct, error) {
	var v map[string]any
	bytes, err := protojson.Marshal(message)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return nil, err
	}
	return structpb.NewStruct(v)
}

func NewMessagingService() protobuftypes.MessagingService {
	return &MessagingService{}
}
