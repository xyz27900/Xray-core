package conf

import (
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/transport/global"
	"github.com/xtls/xray-core/transport/internet"
)

type TransportConfig struct {
	TCPConfig         *TCPConfig          `json:"tcpSettings,omitempty"`
	KCPConfig         *KCPConfig          `json:"kcpSettings,omitempty"`
	WSConfig          *WebSocketConfig    `json:"wsSettings,omitempty"`
	HTTPConfig        *HTTPConfig         `json:"httpSettings,omitempty"`
	DSConfig          *DomainSocketConfig `json:"dsSettings,omitempty"`
	QUICConfig        *QUICConfig         `json:"quicSettings,omitempty"`
	GRPCConfig        *GRPCConfig         `json:"grpcSettings,omitempty"`
	GUNConfig         *GRPCConfig         `json:"gunSettings,omitempty"`
	HTTPUPGRADEConfig *HttpUpgradeConfig  `json:"httpupgradeSettings,omitempty"`
	SplitHTTPConfig   *SplitHTTPConfig    `json:"splithttpSettings,omitempty"`
}

// Build implements Buildable.
func (c *TransportConfig) Build() (*global.Config, error) {
	config := new(global.Config)

	if c.TCPConfig != nil {
		ts, err := c.TCPConfig.Build()
		if err != nil {
			return nil, errors.New("failed to build TCP config").Base(err).AtError()
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "tcp",
			Settings:     serial.ToTypedMessage(ts),
		})
	}

	if c.KCPConfig != nil {
		ts, err := c.KCPConfig.Build()
		if err != nil {
			return nil, errors.New("failed to build mKCP config").Base(err).AtError()
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "mkcp",
			Settings:     serial.ToTypedMessage(ts),
		})
	}

	if c.WSConfig != nil {
		ts, err := c.WSConfig.Build()
		if err != nil {
			return nil, errors.New("failed to build WebSocket config").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "websocket",
			Settings:     serial.ToTypedMessage(ts),
		})
	}

	if c.HTTPConfig != nil {
		ts, err := c.HTTPConfig.Build()
		if err != nil {
			return nil, errors.New("Failed to build HTTP config.").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "http",
			Settings:     serial.ToTypedMessage(ts),
		})
	}

	if c.DSConfig != nil {
		ds, err := c.DSConfig.Build()
		if err != nil {
			return nil, errors.New("Failed to build DomainSocket config.").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "domainsocket",
			Settings:     serial.ToTypedMessage(ds),
		})
	}

	if c.QUICConfig != nil {
		qs, err := c.QUICConfig.Build()
		if err != nil {
			return nil, errors.New("Failed to build QUIC config.").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "quic",
			Settings:     serial.ToTypedMessage(qs),
		})
	}

	if c.GRPCConfig == nil {
		c.GRPCConfig = c.GUNConfig
	}
	if c.GRPCConfig != nil {
		gs, err := c.GRPCConfig.Build()
		if err != nil {
			return nil, errors.New("Failed to build gRPC config.").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "grpc",
			Settings:     serial.ToTypedMessage(gs),
		})
	}

	if c.HTTPUPGRADEConfig != nil {
		hs, err := c.HTTPUPGRADEConfig.Build()
		if err != nil {
			return nil, errors.New("failed to build HttpUpgrade config").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "httpupgrade",
			Settings:     serial.ToTypedMessage(hs),
		})
	}

	if c.SplitHTTPConfig != nil {
		shs, err := c.SplitHTTPConfig.Build()
		if err != nil {
			return nil, errors.New("failed to build SplitHTTP config").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "splithttp",
			Settings:     serial.ToTypedMessage(shs),
		})
	}

	return config, nil
}
