package main

import (
	"flag"
	"time"

	pb "github.com/Amitgb14/gRPC-Examples/heartbeat/proto"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {

	address := flag.String("a", "localhost:8080", "server address:port")
	ser := flag.String("s", "elss", "service name for heart checking")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect %v", err)
	}
	defer conn.Close()

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	c := pb.NewMetricsClient(conn)
	service := &pb.HeartRequest{Service: *ser}

	resp, err := c.HeartStatus(ctx, service)
	if err != nil {
		logrus.Fatalf("response error %v", err)
	}
	logrus.Infof("Request: %s, Response: %s", service, resp.Status)
}
