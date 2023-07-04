package client

type Spec struct {
	Pings []PingConfigBlock `json:"pings"`
}

type PingConfigBlock struct {
	IP   string `json:"ip"`
	NAME string `json:"name"`
}
