codegen:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	mkdir -p ./libs/utils/generated
	protoc -I=./specs/protobuf --go_out=./libs/utils/generated ./specs/protobuf/*.proto --go_opt=paths=source_relative