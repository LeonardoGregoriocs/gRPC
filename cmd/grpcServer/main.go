package main

import (
	"database/sql"
	"net"

	"github.com/LeonardoGregoriocs/gRPC/internal/database"
	"github.com/LeonardoGregoriocs/gRPC/internal/pb"
	"github.com/LeonardoGregoriocs/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()                                // Cria o servidor gRPC.
	pb.RegisterCategoryServiceServer(grpcServer, categoryService) // Conexão entre o servidor (gRPC) e serviço.
	reflection.Register(grpcServer)

	// Abrindo conexão TCP para conseguimos nos comunicar com o gRPC.

	lis, err := net.Listen("tcp", ":50051") // 50051 -> Porta padrão do gRPC.
	if err != nil {
		panic(err)
	}

	// Inicialização do servidor gRPC.

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
