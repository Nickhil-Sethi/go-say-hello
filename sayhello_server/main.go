package main

import "log"
import "net"
import "google.golang.org/grpc"
import pb "github.com/Nickhil-Sethi/go-say-hello/sayhello"

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) WriteLogEvent(ctx *context.Context, logevent *pb.LogEvent) (*pb.LogResponse, error) {
	log.Println(logevent.GetContent())
	return &pb.LogResponse{Status: "success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSayHelloServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
