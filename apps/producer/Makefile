dev:
	npx nodemon --exec "dlv --listen=:2346 --headless --log  --continue --accept-multiclient --api-version=2 debug ./main.go" --ext go -w ./ -w ../../libs/utils/

test: 
	go test -v ./do