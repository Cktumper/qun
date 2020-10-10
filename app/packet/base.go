package packet

type Base struct {
	PacketId  uint   `json:"packet_id"`
	SessionId string `json:"session_id"`
}
