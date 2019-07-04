package tele

import (
	"context"

	tele_config "github.com/temoto/vender/head/tele/config"
	"github.com/temoto/vender/log2"
)

// Tele transport contract:
// - Init fails only with invalid config, ignores network errors
// - Send* deliver (with retries) within timeout or fail; success includes ack from receiver
// - hide "connection" concept from upstream API or errors; transport delivers messages at least once
// - application may start without network available
// - assume worst network quality: packet loss, reorder, duplicates, corruption
type Transporter interface {
	Init(ctx context.Context, log *log2.Log, teleConfig tele_config.Config, onCommand func([]byte) bool, willPayload []byte) error
	SendState(payload []byte) bool
	SendTelemetry(payload []byte) bool
	SendCommandResponse(topicSuffix string, payload []byte) bool
}