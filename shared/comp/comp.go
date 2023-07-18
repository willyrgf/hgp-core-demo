package comp

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/willyrgf/hgp-core-demo/proto/comp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Report struct {
	Message string
	Error   error
}

type Test interface {
	Run(ctx context.Context, e *emptypb.Empty) (*comp.Report, error)
}

type TestPlugin struct {
	plugin.Plugin
	Impl Test
}

func (p *TestPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	comp.RegisterTestServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *TestPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: comp.NewTestClient(c)}, nil
}
