# Golang having gRPC and Rest API at the same time from one proto using grpc-gateway

### How to run:
1. Run `./compile-protos.sh` to generate proto and swagger files.
2. Run `go run server/main.go`
3. Access gRPC server with `grpcurl -plaintext -d '{"name": "John"}' localhost:50051 hello.HelloService/SayHello`
3. Access REST server with `curl -X GET -k http://localhost:50052/api/hello/John`

### How to run separately:
1. For gRPC server run `go run grpc/main.go`
2. For Rest server run `go run rest/main.go`