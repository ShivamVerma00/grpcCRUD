package main

import (
	"fmt"
	"log"
	"net"

	pb "grpcCRUD/protoFile"

	c "grpcCRUD/service"

	"google.golang.org/grpc"
)

const port = ":54321"

func main() {
	fmt.Println("Welcome to Book Management...")

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to Listen...", err)

	}

	//Server Intialisation
	server := grpc.NewServer()

	//Registering server as new grpc server
	pb.RegisterBookManagementServiceServer(server, &c.BookManagementServiceServer{})

	log.Println("Server Listening at", listen.Addr())

	if err := server.Serve(listen); err != nil {
		log.Fatal("Failed to Serve...", err)
	}
}
