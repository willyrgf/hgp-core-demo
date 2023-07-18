package comp

import (
	"github.com/willyrgf/hgp-core-demo/proto/comp"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	client comp.TestClient
}

func (c *GRPCClient) Run(ctx context.Context, e *emptypb.Empty) (*comp.Report, error) {
	resp, err := c.client.Run(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GRPCServer struct {
	Impl Test
	comp.UnimplementedTestServer
}

func (s *GRPCServer) Run(ctx context.Context, e *emptypb.Empty) (*comp.Report, error) {
	return s.Impl.Run(ctx, e)
}
