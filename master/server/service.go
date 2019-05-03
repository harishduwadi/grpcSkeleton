package server

import (
	"context"
	"fmt"

	"github.com/SupplyFrame/gosre/log"

	pb "github.com/harishduwadi/grpcSkeleton/protoFile"
)

// server struct
type Server struct{}

// grpc method defined by the server
func (s *Server) CallServer(ctx context.Context, in *pb.ClientRequest) (*pb.ServerReply, error) {
	log.Infof("Received: %s", in.Message)
	return &pb.ServerReply{Message: fmt.Sprintf("Got: '%s'!\nServer Responds with Hi!", in.Message)}, nil
}
