build:
	find proto -name "*.proto" | xargs protoc -I=. -I=$(PROTO_LIB_PATH) --cato_out=../ --cato_opt=ext_out_dir=../,swagger_path=bms_swagger.json,api_host=localhost