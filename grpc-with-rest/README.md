# Golang having gRPC and Rest API at the same time from one proto using grpc-gateway

### How to run:
1. Generate code from proto with `./compile-protos.sh`
2. Run `go run server/main.go`
3. Access gRPC server with `curl -X GET -k http://localhost:50052/api/hello/John`
3. Access REST server with `grpcurl -plaintext -d '{"name": "John"}' localhost:50051 hello.HelloService/SayHello`

### How to run separately:
1. For gRPC server run `go run grpc/main.go`
2. For Rest server run `go run rest/main.go`

### Swagger
An automatic generation of a swagger file has been added as well, if needed in the future.
