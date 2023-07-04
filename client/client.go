package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
	"github.com/sunil494/cq-source-ping/internal/ping"
)

type Client struct {
	Logger zerolog.Logger
	Spec   *Spec
	Pinger *ping.Client

	Ping PingConfigBlock
}

func (c *Client) ID() string {
	return fmt.Sprintf("ping:%s", c.Ping.NAME)
}

func (c *Client) WithPing(p PingConfigBlock) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("ping", p.IP).Logger()
	newC.Ping = p
	return &newC
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	c, err := ping.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger: logger,
		Spec:   &pluginSpec,
		Pinger: c,
	}, nil
}
