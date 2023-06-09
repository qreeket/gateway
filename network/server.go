package network

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
)

func InitServer() {
	// create a new grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(AuthUnaryInterceptor),
		grpc.StreamInterceptor(AuthStreamInterceptor),
	)

	// register the grpc server for reflection
	reflection.Register(s)

	// @todo register the server

	// get the port number from .env file
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// listen on the port
	if lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port)); err == nil {
		log.Printf("server started on %v\n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("unable to start grpc server: %+v\n", err)
		}
	} else {
		panic(err)
	}
}
