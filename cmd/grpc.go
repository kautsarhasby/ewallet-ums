package cmd

import (
	"kautsarhasby/ewallet-ums/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to listen gRPC port: ", err)
	}

	s := grpc.NewServer()

	logrus.Info("Listening to GRPC :" + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC : ", err)
	}

}
