package main

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/willyrgf/hgp-core-demo/shared"
	"github.com/willyrgf/hgp-core-demo/shared/comp"
	compproto "github.com/willyrgf/hgp-core-demo/proto/comp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Test struct{}

func (Test) Run(ctx context.Context, e *emptypb.Empty) (*compproto.Report, error) {
	return &compproto.Report{Message: "test OK"}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"comp": &comp.TestPlugin{Impl: &Test{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
