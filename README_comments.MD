# 注解大全
## @GORS
必须设置, 基础注解。
代表gors的注解配置的开始，没有`@GORS`开始的都会被忽略。

## @Path
必须设置，路径注解。
- 可以设置参数，`@Path(/api)`
- 如果设置在接口名上，代表为所有path定义基础的path。
- 如果设置在方法上，代表为此方法定义path。
- 可以设置多个`@Path`注解，此时多个path会被拼接成一个。

## HTTP Method注解
必须设置，设置此方法的只能响应指定的Method请求。 Method只能有一个有效，如果设置多个，只有最后一个有效。支持标准库里定义的9种method：
### @GET
>此注解设置此方法为Get请求。
### @HEAD
>此注解设置此方法为Head请求。
### @POST
>此注解设置此方法为Post请求。
### @PUT
>此注解设置此方法为Put请求。
### @PATCH
此注解设置此方法为Patch请求。
### @DELETE
>此注解设置此方法为Delete请求。
### @CONNECT
>此注解设置此方法为Connect请求。
### @OPTIONS
>此注解设置此方法为Options请求。
### @TRACE
>此注解设置此方法为Trace请求。

## Binding 注解
数据绑定注解。
- 当第二个参数是结构体指针时候，可以设置。
- 可以设置多种Binding策略。
### @ReaderBinding
> 配合第二个参数是`io.Reader`类型使用，引用`request.Body`，可以设置`Content-Type`值，比如`@ReaderBinding(image/jpeg)`,不设默认值为`application/octet-stream`
>* grpc定义形式不支持
### @BytesBinding
> 配合第二个参数是`[]byte`类型使用，将`request.Body`以字节形式全部读出，可以设置`Content-Type`值，比如`@ReaderBinding(image/jpeg)`,不设默认值为`application/octet-stream`
>* grpc定义形式不支持
### @StringBinding
> 配合第二个参数是`string`类型使用，将`request.Body`以字节形式全部读出,并转成字符串，可以设置`Content-Type`值，比如`@ReaderBinding(application/javascript)`,不设默认值为`text/plain`
>* grpc定义形式不支持
### @UriBinding
>将对应的Uri参数绑定到结构体指针中。
>接口定义形式，默认的tag是uri,支持[gin的路径参数语法](https://github.com/gin-gonic/gin#parameters-in-path)和[gin的uri绑定策略](https://github.com/gin-gonic/gin#bind-uri)
>* grpc定义形式，默认的tag是json
### @QueryBinding
>将对应的Query参数绑定到结构体指针中。
>* 接口定义形式，默认的tag是form, 支持[gin的只绑定Query策略](https://github.com/gin-gonic/gin#only-bind-query-string)。
>* grpc定义形式，默认的tag是json

### @HeaderBinding
>将对应的Header参数绑定到结构体指针中。
>
>* 接口定义形式，默认的tag是header, 支持[gin的header绑定策略](https://github.com/gin-gonic/gin#bind-header)
>* grpc定义形式，默认的tag是json
 
### @JSONBinding
>将请求的Body以JSON形式解析出来,并绑定到结构体指针中。
>* 接口定义形式，默认的tag是json
>* grpc定义形式，默认的tag是json

### @FormBinding
>将请求的**Query和Body**同时以Form形式解析出来,并绑定到结构体指针中。
>* 接口定义形式，默认的tag是json,支持[gin的form绑定策略](https://github.com/gin-gonic/gin#multiparturlencoded-binding)
>* grpc定义形式，默认的tag是json

### @FormPostBinding
>将请求的**Body**以FormPost形式解析出来,并绑定到结构体指针中。
>* 接口定义形式，默认的tag是json,支持[gin的urlencoded绑定策略](https://github.com/gin-gonic/gin#multiparturlencoded-binding)
>* grpc定义形式，默认的tag是json

### @FormMultipartBinding
>将请求的Body以Multipart形式解析出来并绑定到结构体指针中。
>* 接口定义形式，默认的tag是json,支持文件上传,支持[gin的Multipart绑定策略](https://github.com/gin-gonic/gin#multiparturlencoded-binding)
>* grpc定义形式不支持

### @ProtoBufBinding
>将请求的Body以protobuf形式解析出来并绑定到结构体指针中。
>* 接口定义形式，默认的tag是protobuf
>* grpc定义形式，默认的tag是protobuf

### @ProtoJSONBinding
>将请求的Body以protojson形式解析出来并绑定到结构体指针中。
>* 接口定义形式，默认的tag是json
>* grpc定义形式，默认的tag是json

### @MsgPackBinding
>将请求的Body以msgpack形式解析出来并绑定到结构体指针中。
>* 接口定义形式，默认的tag是codec
>* grpc定义形式，默认的tag是codec

### @XMLBinding
>将请求的Body以XML形式解析出来,并绑定到结构体指针中。
>* 接口定义形式，默认的tag是xml
>* grpc定义形式，默认的tag是xml

### @YAMLBinding
>将请求的Body以YAML形式解析出来并绑定到结构体指针中。
>* 接口定义形式，默认的tag是yaml
>* grpc定义形式，默认的tag是yaml

### @TOMLBinding
>将请求的Body以TOML形式解析出来并绑定到结构体指针中。
>* 接口定义形式，默认的tag是toml
>* grpc定义形式，默认的tag是toml

### @CustomBinding
>预定义的Binding策略不能满足需求，需要使用自定义Binding策略，请求结构体需要实现gors.Binding接口。

## Render 注解
响应渲染注解。
- 方法的一个返回参数的类型不通，相对应的渲染也不同。
- 只能设置一种渲染策略。
### @BytesRender
>字节数组渲染。
>
>只有当方法的第一个返回参数类型为`[]byte`时才能使用，`@BytesRender`将数据以字节数组写入响应的Body中，可以自定义响应`Content-Type`，比如：`@BytesRender(image/png)`代表渲染一张图片，不设默认值为`application/octet-stream`。
>* grpc定义形式不支持

### @StringRender
>字符串渲染。
>
>只有当方法的第一个返回参数类型为`string`时才能使用，`@StringRender`将数据以文本的方式写入响应的Body中，可以自定义响应`Content-Type`，比如`@StringRender(image/xml)`代表渲染xml文件,不设默认值为`text/plain`。
>* grpc定义形式不支持

### @TextRender
>文本渲染。
>
>只有当方法的第一个返回参数类型为`string`时才能使用，`@TextRender`将数据以文本的方式写入响应的Body中，`Content-Type`固定为`text/plain; charset=utf-8`
>* grpc定义形式不支持

### @HTMLRender
>HTML渲染。
>
>只有当方法的第一个返回参数类型为`string`时才能使用，`@HTMLRender`将数据以文本的方式写入响应的Body中，`Content-Type`固定为`text/html; charset=utf-8`
>* grpc定义形式不支持

### @RedirectRender
>重定向渲染。
>
>只有当方法的第一个返回参数类型为`string`时才能使用,此参数应该是个可以重定向的地址位置。
>* grpc定义形式不支持

### @ReaderRender
>数据流渲染，只有当方法的第一个返回参数类型为`io.Reader`时才能使用，`@ReaderRender`将数据流写入响应的Body中，可以自定义响应`Content-Type`，比如：`@ReaderRender(video/mpeg4)`代表渲染一个视频文件,不设默认值为`application/octet-stream`。
>* grpc定义形式不支持

### @JSONRender
>JSON渲染。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@JSONRender`将结构体指针以JSON格式序列化并且写入响应的Body中，`Content-Type`固定为`application/json; charset=utf-8`。参考[gin文档-json](https://github.com/gin-gonic/gin#xml-json-yaml-toml-and-protobuf-rendering)

### @IndentedJSONRender
>漂亮的JSON渲染，在JSON渲染的基础上，加上了缩进和结束行。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@IndentedJSONRender`将结构体指针以JSON格式序列化并且写入响应的Body中，`Content-Type`固定为`application/json; charset=utf-8`。比较耗资源。

### @SecureJSONRender
>安全的的JSON渲染，防止一些代码注入。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@SecureJSONRender`将结构体指针以JSON格式序列化并且写入响应的Body中，`Content-Type`固定为`application/json; charset=utf-8`。[gin文档-SecureJSON](https://github.com/gin-gonic/gin#securejson)

### @JsonpJSONRender
>jsonp方式渲染，jsonp可以跨域请求，如果查询参数callback存在，则将callback添加到响应体中。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用。参考[gin文档-jsonp](https://github.com/gin-gonic/gin#jsonp)
>* grpc定义形式不支持

### @PureJSONRender
>纯净的JSON渲染，普通的JSON渲染，会将一些html特殊字符会被转义，比如`<`转成`\u003c`。`@PureJSONRender`则不会转义。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@PureJSONRender`将结构体指针以JSON格式序列化并且写入响应的Body中，`Content-Type`固定为`application/json; charset=utf-8`。[gin文档-PureJSON](https://github.com/gin-gonic/gin#purejson)

### @AsciiJSONRender
>AsciiJSON渲染，只渲染Ascii字符，非Ascii字符则会被转义。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@AsciiJSONRender`将结构体指针以JSON格式序列化并且写入响应的Body中，`Content-Type`固定为`application/json; charset=utf-8`。[gin文档-AsciiJSON](https://github.com/gin-gonic/gin#asciijson)

### @ProtoJSONRender
>ProtoJSON渲染，对protobuf已Json形式渲染
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@ProtoJSONRender`将结构体指针以JSON格式序列化并且写入响应的Body中，`Content-Type`固定为`application/json; charset=utf-8`。

### @XMLRender
>XML渲染。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@XMLRender`将结构体指针以XML格式序列化并且写入响应的Body中，`Content-Type`固定为`application/xml; charset=utf-8`。参考[gin文档-XMLRender](https://github.com/gin-gonic/gin#xml-json-yaml-toml-and-protobuf-rendering)

### @YAMLRender
>YAML渲染。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@YAMLRender`将结构体指针以YAML格式序列化并且写入响应的Body中，`Content-Type`固定为`application/x-yaml; charset=utf-8`。参考[gin文档-YAMLRender](https://github.com/gin-gonic/gin#xml-json-yaml-toml-and-protobuf-rendering)

### @ProtoBufRender
>ProtoBuf渲染。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@ProtoBufRender`将结构体指针以Protocol Buffer格式序列化并且写入响应的Body中，`Content-Type`固定为`application/x-protobuf`。参考[gin文档-ProtoBufRender](https://github.com/gin-gonic/gin#xml-json-yaml-toml-and-protobuf-rendering)

### @MsgPackRender
>MsgPack渲染。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@MsgPackRender`将结构体指针以MsgPack格式序列化并且写入响应的Body中，`Content-Type`固定为`application/msgpack; charset=utf-8`。

### @TOMLRender
>TOML渲染。
>
>只有当方法的第一个返回参数类型为`结构体指针`时才能使用，`@TOMLRender`将结构体指针以TOML格式序列化并且写入响应的Body中，`Content-Type`固定为`application/toml`。参考[gin文档-TOMLRender](https://github.com/gin-gonic/gin#xml-json-yaml-toml-and-protobuf-rendering)

### @CustomRender
>预定义的Render策略不能满足需求，需要使用自定义Render策略，响应结构体需要实现gors.Render接口。

# Validate
除了gin框架支持的参数校验方案外，gors也提供了另外一种参数校验方式，方法的第一个返回参数类型为`结构体指针`,并且实现了Validator接口，则在数据绑定后，会对数据进行校验。
## Validator接口
```go
type Validator interface {
	Validate() error
}
```