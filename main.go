package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/JigmeTenzinChogyel/bank-bhutan/api"
	db "github.com/JigmeTenzinChogyel/bank-bhutan/db/sqlc"
	"github.com/JigmeTenzinChogyel/bank-bhutan/gapi"
	"github.com/JigmeTenzinChogyel/bank-bhutan/pb"
	"github.com/JigmeTenzinChogyel/bank-bhutan/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// git test

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot conntect to db:", err)
	}

	store := db.NewStore(conn)

	rungRPCServer(config, store)
}

func rungRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBankBhutanServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String()) 
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
