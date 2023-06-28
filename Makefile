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
		--go_opt=module=github.com/go-leo/gors \
		example/api/pb/*.proto \
		internal/pkg/status/*.proto
	@echo "--- protoc generate end ---"

