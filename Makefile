server:
	nodemon --signal SIGTERM --exec "go run ./server.go" --ext go

codegen: 
	nodemon --signal SIGTERM --exec "go run github.com/99designs/gqlgen generate" --ext graphql