package app

type PacketText struct {
	PacketJSON
	Content string `json:"content"`
}

//	构建新的文件包对象
//
//	Author(Wind)
func NewPacketText(content string) *PacketText {
	//	初始化文字包
	p := &PacketText{Content: content}

	//	初始化 JSON 解析器
	p.PacketJSON = PacketJSON{self: p}

	//	返回初始化成功
	return p
}
