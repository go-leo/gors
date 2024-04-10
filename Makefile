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
		example/api/*/*.proto
	@echo "--- protoc generate end ---"

.PHONY: protoc_gors_gen
protoc_gors_gen:
	@echo "--- protoc generate start ---"
	@protoc \
		--proto_path=. \
		--proto_path=cmd/protoc-gen-gors/examples \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--gors_out=. \
		--gors_opt=paths=source_relative \
		--gors_opt=naming=proto \
		--openapi_out=. \
		--openapi_opt=output_mode=source_relative \
		--openapi_opt=naming=proto \
		cmd/protoc-gen-gors/examples/tests/*/*.proto \
		cmd/protoc-gen-gors/examples/google/example/library/*/*.proto
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