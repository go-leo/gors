# 错误码

！！系统错误码列表，由 `gen_error -type=int -doc` 命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：

```json
{
  "code": 100101,
  "message": "Database error"
}
```

上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| Unknown | 100001 | 500 | Internal server error |
| Bind | 100002 | 400 | Error occurred while binding the request body to the struct |
| Validation | 100003 | 400 | Validation failed |
| AccountAuthTypeInvalid | 110001 | 400 | Account AuthType not support |
| UserNotFound | 110002 | 400 | User Not Found |
| UserDisabled | 110003 | 400 | User disabled |

