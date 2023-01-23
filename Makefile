.PHONY: bin tools gen image push integration fmt

bin:
	@ go build -o artifacts/server cmd/server/main.go 

tools:
	@ sudo apt update && sudo apt install -y protobuf-compiler
	@ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

gen:
	@ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/protos/relay/api.proto \
    api/protos/health/api.proto \
    api/protos/beacon/api.proto \

image:
	@ docker build \
		-t troydai/grpcrelay:latest \
		-t troydai/grpcrelay:`git describe --tags` \
		.

fmt:
	@ go fmt ./...

push: image
	@ docker push troydai/grpcrelay:`git describe --tags`

integration:
	@ ./scripts/integration-test.sh
