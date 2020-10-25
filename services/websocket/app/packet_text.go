package app

import "encoding/json"

type PacketText struct {
	Content string `json:"content"`
}

//	构建新的文件包对象
//
//	Author(Wind)
func NewPacketText(content string) *PacketText {
	//	初始化文字包
	p := &PacketText{Content: content}

	//	返回初始化成功
	return p
}

func (p *PacketText) Marshal() []byte {
	b, _ := json.Marshal(p)
	return b
}

func (p *PacketText) Unmarshal(message []byte, v interface{}) error {
	return json.Unmarshal(message, v)
}
