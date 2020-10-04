protobuf-go:
	@protoc \
		--proto_path ./proto \
		--go_opt plugins=grpc \
		--go_opt paths=source_relative \
		--go_out ./proto/gen/go \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_out ./proto/gen/go \
		--swagger_out ./docs/swagger \
		proto/demo/*/*.proto