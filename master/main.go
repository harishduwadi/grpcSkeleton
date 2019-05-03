package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/SupplyFrame/gosre/log"
	"github.com/harishduwadi/grpcSkeleton/master/config"
	"github.com/harishduwadi/grpcSkeleton/master/server"
	pb "github.com/harishduwadi/grpcSkeleton/protoFile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	configFile = "config.xml"
)

func main() {

	xmlConfig, err := parseConfiguration(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	startServices(xmlConfig)
}

func parseConfiguration(configFileLocation string) (*config.Configuration, error) {
	xmlFile, err := ioutil.ReadFile(configFileLocation)
	if err != nil {
		log.Errorf("Error opening file:", err)
		return nil, err
	}
	config := new(config.Configuration)
	err = xml.Unmarshal(xmlFile, &config)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return config, nil

}

func startServices(xmlconfig *config.Configuration) {

	logger := log.New("", xmlconfig.LogFile)
	logger.Start()

	lis, err := net.Listen("tcp", xmlconfig.ListenPort)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return
	}

	creds, err := credentials.NewServerTLSFromFile(xmlconfig.PublicKey, xmlconfig.PrivateKey)
	if err != nil {
		log.Error(err)
		return
	}
	// Below is to connect without ssl
	// s := grpc.NewServer()
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterDbUpdateServer(s, &server.Server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
		logger.Stop()
		return
	}
}
