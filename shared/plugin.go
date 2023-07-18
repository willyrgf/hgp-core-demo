package shared

import (
	"github.com/hashicorp/go-plugin"
	"github.com/willyrgf/hgp-core-demo/shared/comp"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion: 1,
	// FIXME: use UUID
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var PluginMap = map[string]plugin.Plugin{"comp": &comp.TestPlugin{}}
