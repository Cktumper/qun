package packet

//	连接数据包
//	成功 / 失败
type Connect struct {
	SessionId string `json:"session_id"`
	Flag      uint8  `json:"flag"`
}

//	数据包
//
//	Author(Wind)
func (p *Connect) String() string {
	return ""
}
