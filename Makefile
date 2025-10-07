

build-proto:
	protoc -I=. -I="$(PROTO_LIB_PATH)" --cato_out=../ --cato_opt=ext_out_dir=../ ./proto/*/*.proto

build:
	go build -o proto-cato-gen ./cmd