package main

import (
	"context"
	"flag"
	webcli "github.com/codemanufaktur/web-cli/api/generated/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type webCliClient struct {
}

var (
	_          webcli.NewsFeedClient = (*webCliClient)(nil)
	serverAddr                       = flag.String("server_addr", "localhost:9000", "The server address in the format of host:port")
)

func (w webCliClient) ListFeeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*webcli.ListFeedsResponse, error) {
	panic("implement me")
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := webcli.NewNewsFeedClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.ListFeeds(ctx, &emptypb.Empty{})
	log.Printf("got %v responses", len(response.News))
}
