package conf

import (
	"github.com/xtls/xray-core/app/reverse"
	"google.golang.org/protobuf/proto"
)

type BridgeConfig struct {
	Tag    string `json:"tag,omitempty"`
	Domain string `json:"domain,omitempty"`
}

func (c *BridgeConfig) Build() (*reverse.BridgeConfig, error) {
	return &reverse.BridgeConfig{
		Tag:    c.Tag,
		Domain: c.Domain,
	}, nil
}

type PortalConfig struct {
	Tag    string `json:"tag,omitempty"`
	Domain string `json:"domain,omitempty"`
}

func (c *PortalConfig) Build() (*reverse.PortalConfig, error) {
	return &reverse.PortalConfig{
		Tag:    c.Tag,
		Domain: c.Domain,
	}, nil
}

type ReverseConfig struct {
	Bridges []BridgeConfig `json:"bridges,omitempty"`
	Portals []PortalConfig `json:"portals,omitempty"`
}

func (c *ReverseConfig) Build() (proto.Message, error) {
	config := &reverse.Config{}
	for _, bconfig := range c.Bridges {
		b, err := bconfig.Build()
		if err != nil {
			return nil, err
		}
		config.BridgeConfig = append(config.BridgeConfig, b)
	}

	for _, pconfig := range c.Portals {
		p, err := pconfig.Build()
		if err != nil {
			return nil, err
		}
		config.PortalConfig = append(config.PortalConfig, p)
	}

	return config, nil
}
