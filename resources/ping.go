package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/go-ping/ping"
	"github.com/sunil494/cq-source-ping/client"
)

func PingTable() *schema.Table {
	return &schema.Table{
		Name:      "pings",
		Resolver:  fetchPingData,
		Multiplex: client.PingMultiplex,
		Transform: transformers.TransformWithStruct(&ping.Statistics{}),
	}
}

func fetchPingData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	pingConfig := c.Ping
	ping, err := c.Pinger.Ping(pingConfig.IP)
	if err != nil {
		return err
	}
	res <- ping
	return nil
}
