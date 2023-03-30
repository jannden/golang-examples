package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"

	pb "github.com/jannden/golang-examples/todos-with-grpc-postgres-migrate-gorm/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}
var DB *gorm.DB
var err error


func DatabaseConnection() {
	// There are three ways to create a DSN string:
	// Option one:
	/*
	host := "localhost"
	port := "5432"
	dbName := "hello_world"
	dbUser := "postgres"
	password := ""
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			host,
			port,
			dbUser,
			dbName,
			password,
	)
	*/
	
	// Option two - hardcode it all in:
	/*
	dsn := "postgres://postgres:postgres@localhost:5432/hello_world?sslmode=disable"
	*/

	// Option three - use environment variables:
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
			log.Fatal("DATABASE_URL environment variable not set")
	}
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
			log.Fatal("DB connection error: ", err)
	}
	log.Println("DB connection successful")
}

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received: %v", req.GetName())
	todo := &pb.Todo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
	res := DB.Create(&todo)
	if res.RowsAffected == 0 {
			return nil, errors.New("error saving todo")
	}
	return &pb.Todo{
					Id:    todo.Id,
					Name: todo.Name,
					Description: todo.Description,
					Done: false,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("TCP connection failed: %v", err)
	}
	log.Printf("Listening at %v", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}