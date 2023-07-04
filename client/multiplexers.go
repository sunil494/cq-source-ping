package client

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func PingMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for index := range client.Spec.Pings {
		l = append(l, client.WithPing(client.Spec.Pings[index]))
	}
	return l
}
