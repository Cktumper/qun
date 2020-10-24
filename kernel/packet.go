package kernel

//	协议包接口定义
//	该协议包可用于自定义协议
//
//	Author(Wind)
type Packet interface {
	//	将整个包初始化为二进制数据
	//	返回一个二进制数据
	//
	//	Author(Wind)
	Marshal() []byte

	//	将整个包反序列化为一个传入的对象
	//	从 v 参数中返回数据
	//	若出错则返回错误
	//
	//	Author(Wind)
	Unmarshal(message []byte, v interface{}) error
}
