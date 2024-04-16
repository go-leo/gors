package binding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	renderPkg "github.com/go-leo/gors/pkg/render"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
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

func (binding *HttpRuleBinding) Bind(ginCtx *gin.Context, req any) error {
	if _, ok := req.(*emptypb.Empty); ok {
		return nil
	}
	if httpBody, ok := req.(*httpbody.HttpBody); ok {
		return binding.bindHttpBody(ginCtx, httpBody)
	}
	message, ok := req.(proto.Message)
	if !ok {
		return fmt.Errorf("failed convert to proto.Message, %T", req)
	}
	if err := binding.bindParameter(ginCtx, message); err != nil {
		return err
	}
	return binding.bindBody(ginCtx, message)
}

func (binding *HttpRuleBinding) bindHttpBody(ginCtx *gin.Context, body *httpbody.HttpBody) error {
	body.ContentType = ginCtx.ContentType()
	data, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return err
	}
	body.Data = data
	return nil
}

func (binding *HttpRuleBinding) bindParameter(ginCtx *gin.Context, message proto.Message) error {
	parameters := make(map[string]any)
	if err := binding.addPathParameters(ginCtx, parameters); err != nil {
		return err
	}
	binding.addNamedPathParameter(ginCtx, parameters)
	if err := binding.addQueryParameters(ginCtx, parameters); err != nil {
		return err
	}
	if len(parameters) == 0 {
		return nil
	}
	writer := &renderPkg.JSONResponseWriter{Buffer: bytes.Buffer{}}
	if err := render.WriteJSON(writer, parameters); err != nil {
		return err
	}
	return protojson.Unmarshal(writer.Buffer.Bytes(), message)
}

func (binding *HttpRuleBinding) bindBody(ginCtx *gin.Context, message proto.Message) error {
	if binding.Body == nil {
		return nil
	}
	bodyField := binding.Body.Name
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
		if binding.Body.Type == typeString {
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
	if err := protojson.Unmarshal(buffer.Bytes(), bodyMessage); err != nil {
		return err
	}
	proto.Merge(message, bodyMessage)
	return nil
}

func (binding *HttpRuleBinding) addPathParameters(ginCtx *gin.Context, parameters map[string]any) error {
	for _, parameter := range binding.Path {
		val, err := regularValue(parameter.Type, ginCtx.Param(parameter.Name))
		if err != nil {
			return err
		}
		parameters[parameter.Name] = val
	}
	return nil
}

func (binding *HttpRuleBinding) addQueryParameters(ginCtx *gin.Context, parameters map[string]any) error {
	for _, parameter := range binding.Query {
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
			values := ginCtx.QueryArray(parameter.Name)
			queryParameter := make([]any, 0, len(values))
			for _, value := range values {
				val, err := regularValue(parameter.Type, value)
				if err != nil {
					return err
				}
				queryParameter = append(queryParameter, val)
			}
			parameters[name] = parameter
		default:
			val, err := regularValue(parameter.Type, ginCtx.Query(parameter.Name))
			if err != nil {
				return err
			}
			parameters[name] = val
		}
	}
	return nil
}

func (binding *HttpRuleBinding) addNamedPathParameter(ginCtx *gin.Context, parameters map[string]any) {
	if binding.NamedPath == nil {
		return
	}
	args := make([]any, 0, len(binding.NamedPath.Parameters))
	for _, parameterName := range binding.NamedPath.Parameters {
		args = append(args, ginCtx.Param(parameterName))
	}
	value := fmt.Sprintf(binding.NamedPath.Template, args...)
	namedPathParameter := parameters
	namePath := strings.Split(binding.NamedPath.Name, ".")
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
		return strconv.ParseFloat(val, 64)
	case typeInteger:
		return strconv.ParseInt(val, 10, 64)
	case typeBoolean:
		return strconv.ParseBool(val)
	case typeObject:
		return nil, fmt.Errorf("invalid path parameter type: %s, value: %s", typ, typeObject)
	case typeArray:
		return nil, fmt.Errorf("invalid path parameter type: %s, value: %s", typ, typeObject)
	}
	return nil, fmt.Errorf("unsurpport type: %s", typ)
}

var (
	typeString  = "string"
	typeNumber  = "number"
	typeInteger = "integer"
	typeBoolean = "boolean"
	typeObject  = "object"
	typeArray   = "array"
)
