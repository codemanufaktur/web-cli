package webcliclient

import (
	"context"
	webcli "github.com/codemanufaktur/web-cli/api/generated/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type WebCliClient struct {
	client webcli.NewsFeedClient
}

var (
	_ webcli.NewsFeedClient = (*WebCliClient)(nil)
)

func NewClient(serverAddress string) *WebCliClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := webcli.NewNewsFeedClient(conn)
	//TODO who closes the grpc conn?

	c := &WebCliClient{client: client}
	return c
}

func (w WebCliClient) ListFeeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*webcli.ListFeedsResponse, error) {

	response, err := w.client.ListFeeds(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return response, nil
}
