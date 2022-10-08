#!/usr/bin/env sh

protoc \
  --proto_path=. \
  --go_out=. \
  --go_opt=module=github.com/go-leo/example \
  --go-httpx_out=. \
  --go-httpx_opt=module=github.com/go-leo/example \
  api/*/*.proto
