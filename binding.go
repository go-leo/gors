package gors

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	internalbinding "github.com/go-leo/gors/internal/pkg/binding"
	"github.com/go-leo/gox/stringx"
	"io"
	"net/http"
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

//
//func AutoBinding(ctx context.Context, req any, _ string) error {
//	c := FromContext(ctx)
//	switch {
//	case strings.Contains(c.ContentType(), annotation.FormContentType):
//		return FormBinding(ctx, req, annotation.FormContentType)
//	case strings.Contains(c.ContentType(), annotation.FormMultipartContentType):
//		return FormMultipartBinding(ctx, req, annotation.FormMultipartContentType)
//	case strings.Contains(c.ContentType(), annotation.JSONContentType):
//		return JSONBinding(ctx, req, annotation.JSONContentType)
//	case strings.Contains(c.ContentType(), annotation.XMLContentType):
//		return XMLBinding(ctx, req, annotation.XMLContentType)
//	case strings.Contains(c.ContentType(), annotation.XML2ContentType):
//		return XMLBinding(ctx, req, annotation.XML2ContentType)
//	case strings.Contains(c.ContentType(), annotation.ProtoBufContentType):
//		return ProtoBufBinding(ctx, req, annotation.ProtoBufContentType)
//	case strings.Contains(c.ContentType(), annotation.MsgPackContentType):
//		return MsgPackBinding(ctx, req, annotation.MsgPackContentType)
//	case strings.Contains(c.ContentType(), annotation.MsgPack2ContentType):
//		return MsgPackBinding(ctx, req, annotation.MsgPack2ContentType)
//	case strings.Contains(c.ContentType(), annotation.YAMLContentType):
//		return YAMLBinding(ctx, req, annotation.YAMLContentType)
//	case strings.Contains(c.ContentType(), annotation.TOMLContentType):
//		return TOMLBinding(ctx, req, annotation.TOMLContentType)
//	default:
//		customBinding, ok := req.(Binding)
//		if !ok {
//			return nil
//		}
//		return customBinding.Bind(ctx)
//	}
//}
