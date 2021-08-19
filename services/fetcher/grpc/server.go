package grpc

import (
	"context"

	pb "github.com/hatena/Hatena-Intern-2021/services/fetcher/pb/fetcher"
	"github.com/hatena/Hatena-Intern-2021/services/fetcher/fetcher"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb.UnimplementedFetchererServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb.FetcherRequest) (*pb.FetcherReply, error) {
	title, err := fetcher.Fetcher(ctx, in.Url)
	if err != nil {
		return nil, err
	}
	return &pb.FetcherReply{Title: title}, nil
}
