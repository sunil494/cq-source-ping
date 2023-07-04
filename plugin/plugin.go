package plugin

import (
	"github.com/sunil494/cq-source-ping/client"
	"github.com/sunil494/cq-source-ping/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"sunil494-ping",
		Version,
		schema.Tables{
			resources.PingTable(),
		},
		client.New,
	)
}
