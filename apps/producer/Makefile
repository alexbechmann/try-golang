dev:
	npx nodemon --signal SIGTERM --exec "dlv --listen=:2346 --headless --log  --continue --accept-multiclient --api-version=2 debug ./main.go --output ./debug/__debug__" --ext go -w ./ -w ../../libs/utils/ 

test: 
	go test -v ./do