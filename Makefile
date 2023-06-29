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
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/go-leo/gors \
		--go-gors_out=. \
		--go-gors_opt=module=github.com/go-leo/gors \
		--go-gors_opt=path_to_lower=true \
		--go-gors_opt=restful=true \
		example/api/*/*.proto
	@echo "--- protoc generate end ---"