package main

import (
	"context"
	"flag"
	"fmt"
	webcli "github.com/codemanufaktur/web-cli/api/generated/proto"
	"github.com/codemanufaktur/web-cli/cmd"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

type webCliServer struct {
	webcli.UnimplementedNewsFeedServer
}

var (
	_    webcli.NewsFeedServer = (*webCliServer)(nil)
	port                       = flag.Int("port", 9000, "server port")
)

func (w webCliServer) ListFeeds(ctx context.Context, empty *emptypb.Empty) (*webcli.ListFeedsResponse, error) {

	news := cmd.GetNewsList(5)

	var realNews []*webcli.News

	for i := 0; i < len(news.News); i++ {

		news, err := mapNews(news.News[i])
		if err != nil {
			return nil, err
		}

		realNews = append(realNews, news)
	}
	return &webcli.ListFeedsResponse{News: realNews}, nil
}

func mapNews(news cmd.News) (response *webcli.News, err error) {

	ref := "2006-01-02T15:04:05-07:00"
	t, err := time.Parse(ref, news.Date)
	if err != nil {
		return nil, err
	}

	return &webcli.News{
		ID:          news.ID,
		Title:       news.Title,
		Description: news.Description,
		Date:        &timestamppb.Timestamp{Seconds: t.Unix()},
	}, nil
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
