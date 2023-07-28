dev:
	npx nodemon --signal SIGTERM --exec "dlv --listen=:2345 --headless --log  --continue --accept-multiclient --api-version=2 debug ./main.go --output ./debug/__debug__" --ext go -w ./ -w ../../libs/utils/ 

test: mockgen
	go test -v ./do

mockgen:
	# go install go.uber.org/mock/mockgen@latest
	mockgen -package=testutils -source=./kafkahandler/kafkahandler.go > ./testutils/kafkahandler_mocks.go 
	mockgen -package=testutils  -source=./store/store.go > ./testutils/store_mocks.go
