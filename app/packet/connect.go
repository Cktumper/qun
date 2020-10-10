package packet

//	连接数据包
//	成功 / 失败
type Connect struct {
	Base
	Flag uint8 `json:"flag"`
}
