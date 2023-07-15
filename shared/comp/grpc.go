package comp

import (
	"github.com/willyrgf/hgp-core-demo/proto/comp"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	client comp.TestClient
}

func (c *GRPCClient) Run(ctx context.Context) (*comp.Report, error) {
	resp, err := c.client.Run(ctx, &comp.Empty{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GRPCServer struct {
	Impl Test
	comp.UnimplementedTestServer
}

func (s *GRPCServer) Run(ctx context.Context, e *comp.Empty) (*comp.Report, error) {
	return s.Impl.Run(ctx, e)
}
