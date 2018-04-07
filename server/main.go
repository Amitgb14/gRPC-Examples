package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/Amitgb14/gRPC-Examples/heartbeat/proto"
)

// server is used to implement .
type server struct{}

func (s *server) HeartStatus(ctx context.Context, in *pb.HeartRequest) (*pb.HeartResponse, error) {
	return &pb.HeartResponse{Status: "green"}, nil

}

func main() {
	port := flag.Int("p", 8080, "port to listen.")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}
	logrus.Infof("Server Listening on %d", *port)

	s := grpc.NewServer()

	pb.RegisterMetricsServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve %v", err)
	}
}
