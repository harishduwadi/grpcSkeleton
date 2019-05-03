package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/SupplyFrame/gosre/log"

	"github.com/harishduwadi/grpcSkeleton/agent/config"
	pb "github.com/harishduwadi/grpcSkeleton/protoFile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func startServices(xmlconfig *config.Configuration) {
	// Set up a connection to the server.

	logger := log.New("", xmlconfig.LogFile)
	logger.Start()

	creds, err := credentials.NewClientTLSFromFile(xmlconfig.PublicKeyLocation, "")
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Dialing to the server")
	// Below is to connect without ssl
	// conn, err := grpc.Dial(xmlconfig.ListenPort, grpc.WithInsecure())
	conn, err := grpc.Dial(xmlconfig.ListenPort, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	log.Info("Connection to the server has been established!")
	c := pb.NewDbUpdateClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CallServer(ctx, &pb.ClientRequest{Message: "Client Says Hello!"})
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return
	}
	// Returned struct from the server
	log.Infof("%s", r.Message)
	logger.Stop()
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
