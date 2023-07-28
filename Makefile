codegen:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	mkdir -p ./libs/utils/protos
	protoc -I=./specs/protobuf --go_out=./libs/utils/protos ./specs/protobuf/*.proto --go_opt=paths=source_relative

test:
	cd apps/example && go test -v ./...
	cd libs/utils && go test -v ./...
	cd apps/consumer && make test