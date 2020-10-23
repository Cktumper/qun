package kernel

import "errors"

var (
	//	房间
	ErrRoomIsFull          = errors.New("该房间已人满为患")
	ErrRoomIsEmpty         = errors.New("该房间已经没人了")
	ErrRoomNotFoundSession = errors.New("查无此人")

	//	房间容器
	ErrContainerNotFoundRoom = errors.New("查无此房间")
)
