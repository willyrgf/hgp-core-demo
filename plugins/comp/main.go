package main

import (
	"context"

	"github.com/willyrgf/hgp-core-demo/proto/comp"
)

type Test struct{}

func (Test) Run(ctx context.Context) (*comp.Report, error) {
	return &comp.Report{Message: "test OK"}, nil
}

func main() {
}
