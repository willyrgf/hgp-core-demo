// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/willyrgf/hgp-core-demo/shared"
	"github.com/willyrgf/hgp-core-demo/shared/comp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func run() error {
	// We're a host. Start by launching the plugin process.
	pluginName := os.Getenv("PLUGIN")
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", "./"+pluginName),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(pluginName)
	if err != nil {
		return err
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	test := raw.(comp.Test)
	report, err := test.Run(context.TODO(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Printf("Report:\n%+v\n", report)
	return nil
}

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(io.Discard)

	err := run()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
