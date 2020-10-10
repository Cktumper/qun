package packet

type Chat struct {
	Base
	Message string `json:"message"`
}
