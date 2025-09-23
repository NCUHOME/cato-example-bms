

build-proto:
	protoc -I=. -I="$(proto_root_path)" --cato_out=../ ./proto/*/*.proto

build:
	go build -o proto-cato-gen ./cmd