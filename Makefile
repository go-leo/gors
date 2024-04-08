.PHONY: go_gen
go_gen:
	@echo "--- go generate start ---"
	@go generate ./...
	@echo "--- go generate end ---"

.PHONY: protoc_gen
protoc_gen:
	@echo "--- protoc generate start ---"
	@protoc \
		--proto_path=. \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--go-gors_out=. \
		--go-gors_opt=paths=source_relative \
		--jsonschema_out=. \
		--jsonschema_opt=paths=source_relative \
		--gorsopenapi_out=. \
		--gorsopenapi_opt=paths=source_relative \
		example/api/*/*.proto
	@echo "--- protoc generate end ---"

protoc_pkg_gen:
	@echo "--- protoc generate start ---"
	@protoc \
		--proto_path=. \
		--go_out=. \
		--go_opt=module=github.com/go-leo/gors \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/go-leo/gors \
		--go-gors_out=. \
		--go-gors_opt=module=github.com/go-leo/gors \
		internal/pkg/*/*.proto
	@echo "--- protoc generate end ---"