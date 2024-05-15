package binding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/render"
	renderPkg "github.com/go-leo/gors/pkg/render"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// PathRule /{message_id} => /:message_id
type PathRule struct {
	// Name message_id
	Name string
	// Type string,integer,number,boolean
	Type string
}

// NamedPathRule /{book.name=shelves/*/books/*} => /shelves/{shelf}/books/{book}
type NamedPathRule struct {
	// Name book.name
	Name string
	// Parameters shelf,book
	Parameters []string
	// Template shelves/%s/books/%s
	Template string
}

// QueryRule ?message_id=1
type QueryRule struct {
	// Name message_id
	Name string
	// Type string,integer,number,boolean,array
	Type string
	// ItemType
	ItemType string
}

// BodyRule body
type BodyRule struct {
	// Name * or xxx
	Name string
	// Type string,integer,number,boolean,array,object
	Type string
}

type HttpRuleBinding struct {
	Path      []*PathRule
	NamedPath *NamedPathRule
	Query     []*QueryRule
	Body      *BodyRule
}

func (b *HttpRuleBinding) Bind(ginCtx *gin.Context, req any) error {
	if _, ok := req.(*emptypb.Empty); ok {
		return nil
	}
	if httpBody, ok := req.(*httpbody.HttpBody); ok {
		return b.bindHttpBody(ginCtx, httpBody)
	}
	if request, ok := req.(*rpchttp.HttpRequest); ok {
		return b.bindHttpRequest(ginCtx, request)
	}
	message, ok := req.(proto.Message)
	if !ok {
		return fmt.Errorf("failed convert to proto.Message, %T", req)
	}
	if err := b.bindParameter(ginCtx, message); err != nil {
		return err
	}
	return b.bindBody(ginCtx, message)
}

func (b *HttpRuleBinding) bindHttpBody(ginCtx *gin.Context, body *httpbody.HttpBody) error {
	body.ContentType = ginCtx.ContentType()
	data, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return err
	}
	body.Data = data
	return nil
}

func (b *HttpRuleBinding) bindHttpRequest(ctx *gin.Context, request *rpchttp.HttpRequest) error {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	request.Body = data
	request.Method = ctx.Request.Method
	request.Uri = ctx.Request.RequestURI
	for key := range ctx.Request.Header {
		request.Headers = append(request.Headers, &rpchttp.HttpHeader{
			Key:   http.CanonicalHeaderKey(key),
			Value: ctx.Request.Header.Get(key),
		})
	}
	return nil
}

func (b *HttpRuleBinding) bindParameter(ginCtx *gin.Context, message proto.Message) error {
	parameters := make(map[string]any)
	if err := b.addPathParameters(ginCtx, parameters); err != nil {
		return err
	}
	b.addNamedPathParameter(ginCtx, parameters)
	if err := b.addQueryParameters(ginCtx, parameters); err != nil {
		return err
	}
	if len(parameters) == 0 {
		return nil
	}
	writer := &renderPkg.JSONResponseWriter{Buffer: bytes.Buffer{}}
	if err := render.WriteJSON(writer, parameters); err != nil {
		return err
	}
	return binding.JSON.BindBody(writer.Buffer.Bytes(), message)
}

func (b *HttpRuleBinding) bindBody(ginCtx *gin.Context, message proto.Message) error {
	if b.Body == nil {
		return nil
	}
	bodyField := b.Body.Name
	if len(bodyField) == 0 {
		return nil
	}
	bodyData, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}
	if bodyField == "*" {
		buffer.Write(bodyData)
	} else {
		buffer.WriteString(`{"` + bodyField + `":`)
		if b.Body.Type == typeString {
			data, err := json.Marshal(string(bodyData))
			if err != nil {
				return err
			}
			buffer.Write(data)
		} else {
			buffer.Write(bodyData)
		}
		buffer.WriteString("}")
	}
	var bodyMessage proto.Message
	messageRef := message.ProtoReflect()
	if !messageRef.IsValid() {
		bodyMessage = messageRef.Type().Zero().Interface()
	} else {
		dst := messageRef.New()
		bodyMessage = dst.Interface()
	}
	if err := binding.JSON.BindBody(buffer.Bytes(), bodyMessage); err != nil {
		return err
	}
	proto.Merge(message, bodyMessage)
	return nil
}

func (b *HttpRuleBinding) addPathParameters(ginCtx *gin.Context, parameters map[string]any) error {
	for _, parameter := range b.Path {
		val, err := regularValue(parameter.Type, ginCtx.Param(parameter.Name))
		if err != nil {
			return err
		}
		parameters[parameter.Name] = val
	}
	return nil
}

func (b *HttpRuleBinding) addQueryParameters(ginCtx *gin.Context, parameters map[string]any) error {
	for _, parameter := range b.Query {
		queryParameters := parameters
		namePath := strings.Split(parameter.Name, ".")
		for _, nameSeg := range namePath[:len(namePath)-1] {
			if _, ok := queryParameters[nameSeg]; !ok {
				queryParameters[nameSeg] = make(map[string]any)
			}
		}
		name := namePath[len(namePath)-1]
		switch parameter.Type {
		case typeArray:
			values, ok := ginCtx.GetQueryArray(parameter.Name)
			if !ok {
				return nil
			}
			queryParameter := make([]any, 0, len(values))
			for _, value := range values {
				val, err := regularValue(parameter.ItemType, value)
				if err != nil {
					return err
				}
				queryParameter = append(queryParameter, val)
			}
			parameters[name] = queryParameter
		default:
			value, ok := ginCtx.GetQuery(parameter.Name)
			if !ok {
				return nil
			}
			val, err := regularValue(parameter.Type, value)
			if err != nil {
				return err
			}
			parameters[name] = val
		}
	}
	return nil
}

func (b *HttpRuleBinding) addNamedPathParameter(ginCtx *gin.Context, parameters map[string]any) {
	if b.NamedPath == nil {
		return
	}
	args := make([]any, 0, len(b.NamedPath.Parameters))
	for _, parameterName := range b.NamedPath.Parameters {
		args = append(args, ginCtx.Param(parameterName))
	}
	value := fmt.Sprintf(b.NamedPath.Template, args...)
	namedPathParameter := parameters
	namePath := strings.Split(b.NamedPath.Name, ".")
	for _, nameSeg := range namePath[:len(namePath)-1] {
		if _, ok := namedPathParameter[nameSeg]; !ok {
			subParameter := make(map[string]any)
			namedPathParameter[nameSeg] = subParameter
			namedPathParameter = subParameter
		}
	}
	namedPathParameter[namePath[len(namePath)-1]] = value
}

func regularValue(typ string, val string) (any, error) {
	switch typ {
	case typeString:
		return val, nil
	case typeNumber:
		if val == "" {
			return float64(0), nil
		}
		return strconv.ParseFloat(val, 64)
	case typeInteger:
		if val == "" {
			return int64(0), nil
		}
		return strconv.ParseInt(val, 10, 64)
	case typeBoolean:
		if val == "" {
			return false, nil
		}
		return strconv.ParseBool(val)
	case typeArray, typeObject:
		return nil, fmt.Errorf("invalid parameter type: %s, value: %s", typ, typeObject)
	default:
		return nil, fmt.Errorf("unsurpport type: %s", typ)
	}
}

var (
	typeString  = "string"
	typeNumber  = "number"
	typeInteger = "integer"
	typeBoolean = "boolean"
	typeObject  = "object"
	typeArray   = "array"
)
