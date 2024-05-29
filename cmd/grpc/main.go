package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/andrefsilveira1/grpc/internal/database"
	"github.com/andrefsilveira1/grpc/internal/pb"
	"github.com/andrefsilveira1/grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	var port = ":50051"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server starting on port %s", port)
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
