package chat

import (
	"log"
	"sync"
)

//	聊天房间结构
//
//	Author(Wind)
type Room struct {
	mu sync.Mutex

	RoomId       int
	Name         string
	Limit, Total int

	//	房间里的会话
	sessions []Session
}

//	构建一个新的房间
//
//	Author(Wind)
func NewRoom(roomId int, name string, limit int) *Room {
	return &Room{RoomId: roomId, Name: name, Limit: limit}
}

//	加入房间
//
//	Author(Wind)
func (p *Room) Join(session Session) error {
	//	并发安全 ，此操作需要加锁
	p.mu.Lock()
	defer p.mu.Unlock()

	//	判定房间是否满员了
	if p.isFull() {
		return ErrRoomIsFull
	}

	//	若房间未满员加入至会话中
	p.sessions = append(p.sessions, session)

	//	增加人数
	p.Total++

	//	执行接收消息
	go func() {
		for {
			//	接收消息
			packet, reply, err := session.GetConnection().Receiver()

			//	如果出错则退出消息循环
			if err != nil {
				log.Printf("%s 关闭了连接", session.GetUserInformation().GetNickname())
				break
			}

			//	判定是否回信，返回 true 则回信
			if reply {
				_ = p.Send(packet)
				go log.Printf("%s 发送了消息 %s", session.GetUserInformation().GetNickname(), string(packet.Marshal()))
			}
		}
	}()

	//	返回成功
	return nil
}

//	离开房间
//
//	Author(Wind)
func (p *Room) Leave(session Session) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	//	判定房间是否还有人
	if p.Total <= 0 {
		return ErrRoomIsEmpty
	}

	//	如果房间还有人则进行离开操作
	//	先查找人的游标
	index := p.SearchPosition(session)

	//	如果游标返回小于0，则表示查无此人
	if index < 0 {
		return ErrRoomNotFoundSession
	}

	//	删除游标下的人
	//	这里对切片下的用户进行分割
	//	如果在游标首位则去头，如果在游标末位则去尾，如果在游标中间就掐头去尾
	switch {
	case index == 0:
		p.sessions = p.sessions[1:]
	case index == len(p.sessions):
		p.sessions = p.sessions[:len(p.sessions)-1]
	default:
		p.sessions = append(p.sessions[:index-1], p.sessions[index+1:]...)
	}

	//	减去计数器
	p.Total--

	//	关闭该 Session 连接
	session.GetConnection().Close()

	//	返回成功
	return nil
}

//	向房间里的所有人群发消息
//
//	Author(Wind)
func (p *Room) Send(packet Packet) error {
	//	循环遍历会话，并发送消息
	for _, session := range p.sessions {
		//	TODO: 此处直接使用 Go程并不好，花时间需要优化它
		go session.GetConnection().Send(packet)
	}

	//	返回成功
	return nil
}

//	查找某一个人
//
//	Author(Wind)
func (p *Room) SearchPosition(session Session) int {
	//	查找某一个人
	for row, s := range p.sessions {
		if s.GetUserInformation().GetUserId() == session.GetUserInformation().GetUserId() {
			return row
		}
	}

	//	如果未找到这个人，则返回 -1
	return -1
}

//	判定当前房间是否满员了
//	如果满员了，则返回 True ,否则返回 False
//
//	Author(Wind)
func (p *Room) isFull() bool {
	return p.Total >= p.Limit
}
