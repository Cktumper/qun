package app

import (
	"encoding/json"
	"peon.top/qun/chat"
)

//	数据包解析方式
//	这里使用 JSON
type PacketJSON struct {
	self chat.Packet
}

//	序列化
//
//	Author(Wind)
func (p *PacketJSON) Marshal() []byte {
	byt, _ := json.Marshal(p.self)
	return byt
}

//	解析 JSON 包
//
//	Author(Wind)
func (p *PacketJSON) Unmarshal(message []byte, v interface{}) error {
	return json.Unmarshal(message, v)
}
