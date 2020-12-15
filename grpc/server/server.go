package main

import (
	"context"
	"flag"
	"fmt"
	webcli "github.com/codemanufaktur/web-cli/api/generated/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type webCliServer struct {
	webcli.UnimplementedNewsFeedServer
}

var (
	_    webcli.NewsFeedServer = (*webCliServer)(nil)
	port                       = flag.Int("port", 9000, "server port")
)

func (w webCliServer) ListFeeds(ctx context.Context, empty *emptypb.Empty) (*webcli.ListFeedsResponse, error) {
	panic("implement me")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	webcli.RegisterNewsFeedServer(grpcServer, &webCliServer{})
	grpcServer.Serve(lis)
}
