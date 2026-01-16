package xctrl

import (
	client "github.com/xswitch-cn/proto/xctrl/client"

	"encoding/json"
)

type XNativeJSRequestData struct {
	Command string          `json:"command,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

type XNativeJSRequest struct {
	CtrlUuid string                `json:"ctrl_uuid,omitempty"`
	Data     *XNativeJSRequestData `json:"data,omitempty"`
}

type XNativeJSResponse struct {
	Code     int32  `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
	NodeUuid string `json:"node_uuid,omitempty"`
	// optional
	Seq  string           `json:"seq,omitempty"`
	Data *json.RawMessage `json:"data,omitempty"`
}

var service *XNodeService

func Service() XNodeService {
	return *service
}

func SetService(s *XNodeService) {
	service = s
}

// build default NATS Node Address
func (c *ChannelEvent) WithAddress() client.CallOption {
	prefix := c.PNATSToPrefix + c.PNATSTenantID
	if prefix != "" {
		prefix += "."
	}
	subject := prefix + "cn.xswitch.node." + c.GetNodeUuid()
	return client.WithAddress(subject)
}

func (c *ChannelEvent) SetToPrefix(prefix string) {
	c.PNATSToPrefix = prefix
}

func (c *ChannelEvent) SetTenantID(tenant string) {
	c.PNATSTenantID = tenant
}
