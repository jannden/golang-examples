# Golang, gRPC, Postgres with Gorm and golang-migrate
## Todo app

This is a simple todo app written in Go, using gRPC and Postgres.

### How to run:
1. Set the environment variable `DATABASE_URL` to your Postgres connection string
2. Run the migrations with `migrate -database ${DATABASE_URL} -path migrations up`
3. Generate the gRPC code with `./compile-protos.sh`
4. Run `go run server/main.go`
5. Run `go run client/main.go`