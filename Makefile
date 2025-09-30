

build-proto:
	protoc -I=. -I="$(PROTO_LIB_PATH)" --cato_out=../ ./proto/*/*.proto

build:
	go build -o proto-cato-gen ./cmd