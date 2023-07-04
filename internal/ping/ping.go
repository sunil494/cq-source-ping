package ping

import (
	"github.com/go-ping/ping"
)

type Client struct {
}

type Option func(*Client)

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}
	return c, nil
}

func (c *Client) Ping(IP string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(IP)
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	return stats, nil
}
