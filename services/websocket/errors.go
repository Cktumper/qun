package websocket

import "errors"

var (
	ErrNotSupportTLS = errors.New("还未支持TLS，请等待后续更新")
)
