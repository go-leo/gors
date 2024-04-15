package gors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	internalbinding "github.com/go-leo/gors/internal/pkg/binding"
	"github.com/go-leo/gox/stringx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func RequestBind(ctx context.Context, req any, tag string, bindings ...func(ctx context.Context, req any, tag string) error) error {
	for _, fn := range bindings {
		if err := fn(ctx, req, tag); err != nil {
			return fmt.Errorf("failed to bind request, %w", err)
		}
	}
	if err := Validate(req); err != nil {
		return fmt.Errorf("failed to validate request, %w", err)
	}
	return nil
}

func StringBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	reqPtr, ok := req.(*string)
	if !ok {
		return fmt.Errorf("%T not converted to *string", req)
	}
	*reqPtr = string(body)
	return nil
}

func BytesBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	reqPtr, ok := req.(*[]byte)
	if !ok {
		return fmt.Errorf("%T not converted to *[]byte", req)
	}
	*reqPtr = body
	return nil
}

func ReaderBinding(ctx context.Context, req any, _ string) error {
	c := FromContext(ctx)
	reqPtr, ok := req.(*io.Reader)
	if !ok {
		return fmt.Errorf("%T not converted to *io.Reader", req)
	}
	*reqPtr = c.Request.Body
	return nil
}

func UriBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	if stringx.IsBlank(tag) {
		return c.ShouldBindUri(req)
	}
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return binding.MapFormWithTag(req, m, tag)
}

func QueryBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.Query)
	}
	return binding.MapFormWithTag(req, c.Request.URL.Query(), tag)
}

func HeaderBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.Header)
	}
	return binding.MapFormWithTag(req, c.Request.Header, tag)
}

func FormBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.Form)
	}
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	const defaultMemory = 32 << 20
	if err := c.Request.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return binding.MapFormWithTag(req, c.Request.Form, tag)
}

func FormPostBinding(ctx context.Context, req any, tag string) error {
	c := FromContext(ctx)
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.FormPost)
	}
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	return binding.MapFormWithTag(req, c.Request.PostForm, tag)
}

func FormMultipartBinding(ctx context.Context, req any, tag string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.FormMultipart)
}

func JSONBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.JSON)
}

func ProtoJSONBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, internalbinding.ProtoJSON)
}

func XMLBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.XML)
}

func ProtoBufBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.ProtoBuf)
}

func MsgPackBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.MsgPack)
}

func YAMLBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.YAML)
}

func TOMLBinding(ctx context.Context, req any, _ string) error {
	return FromContext(ctx).ShouldBindWith(req, binding.TOML)
}

func CustomBinding(ctx context.Context, req any, _ string) error {
	customBinding, ok := req.(Binding)
	if !ok {
		return nil
	}
	return customBinding.Bind(ctx)
}

// PathParameter /{message_id} => /:message_id
type PathParameter struct {
	// Name message_id
	Name string
	// Type string,integer,number,boolean
	Type string
}

// NamedPathParameter /{book.name=shelves/*/books/*} => /shelves/{shelf}/books/{book}
type NamedPathParameter struct {
	// Name book.name
	Name string
	// Parameters shelf,book
	Parameters []string
	// Template shelves/%s/books/%s
	Template string
}

// QueryParameter ?message_id=1
type QueryParameter struct {
	// Name message_id
	Name string
	// Type string,integer,number,boolean,array
	Type string
	// ItemType
	ItemType string
}

// BodyParameter body
type BodyParameter struct {
	// Name * or xxx
	Name string
	// Type string,integer,number,boolean,array,object
	Type string
}

type Payload struct {
	Path      []*PathParameter
	NamedPath *NamedPathParameter
	Query     []*QueryParameter
	Body      *BodyParameter
}

func PayloadBinding(payload *Payload) func(ctx context.Context, req any, tag string) error {
	return func(ctx context.Context, req any, tag string) error {
		ginCtx := FromContext(ctx)
		message, ok := req.(proto.Message)
		if !ok {
			return fmt.Errorf("failed convert to proto.Message, %T", req)
		}
		parameters := make(map[string]any)

		if err := addPathParameters(payload, ginCtx, parameters); err != nil {
			return err
		}

		addNamedPathParameter(payload, ginCtx, parameters)

		if err := addQueryParameters(payload, ginCtx, parameters); err != nil {
			return err
		}

		parameterJson, err := json.Marshal(parameters)
		if err != nil {
			return err
		}
		if err := protojson.Unmarshal(parameterJson, message); err != nil {
			return err
		}

		if payload.Body != nil {
			if payload.Body.Name != "*" {
				return nil
			}
		}

		return nil
	}
}

func addQueryParameters(payload *Payload, ginCtx *gin.Context, parameters map[string]any) error {
	for _, parameter := range payload.Query {
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

func addNamedPathParameter(payload *Payload, ginCtx *gin.Context, parameters map[string]any) {
	if payload.NamedPath == nil {
		return
	}
	args := make([]any, 0, len(payload.NamedPath.Parameters))
	for _, parameterName := range payload.NamedPath.Parameters {
		args = append(args, ginCtx.Param(parameterName))
	}
	value := fmt.Sprintf(payload.NamedPath.Template, args...)
	namedPathParameter := parameters
	namePath := strings.Split(payload.NamedPath.Name, ".")
	for _, nameSeg := range namePath[:len(namePath)-1] {
		if _, ok := namedPathParameter[nameSeg]; !ok {
			subParameter := make(map[string]any)
			namedPathParameter[nameSeg] = subParameter
			namedPathParameter = subParameter
		}
	}
	namedPathParameter[namePath[len(namePath)-1]] = value
}

func addPathParameters(payload *Payload, ginCtx *gin.Context, parameters map[string]any) error {
	for _, parameter := range payload.Path {
		val, err := regularValue(parameter.Type, ginCtx.Param(parameter.Name))
		if err != nil {
			return err
		}
		parameters[parameter.Name] = val
	}
	return nil
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

	formatDate     = "date"
	formatDateTime = "date-time"
	formatEnum     = "enum"
	formatBytes    = "bytes"
)
