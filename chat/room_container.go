package chat

import "sync"

type RoomContainer struct {
	mu sync.Mutex

	rooms map[int]*Room
	total int
}

var (
	RoomContainerInstance     *RoomContainer
	roomContainerOnceInstance sync.Once
)

//	获取房间容器单例
//	房间容器只能初始化一次
//
// Author(Wind)
func GetRoomContainer() *RoomContainer {
	roomContainerOnceInstance.Do(func() {
		RoomContainerInstance = &RoomContainer{rooms: make(map[int]*Room)}
	})

	return RoomContainerInstance
}

//	插入一个房间到容器中
//
//	Author(Wind)
func (p *RoomContainer) Insert(room *Room) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	//	将房间对象加入至房间列表中
	p.rooms[room.RoomId] = room

	//	增加房间数
	p.total++

	//	返回成功
	return nil
}

//	删除容器中的一个房间
//
//	Author(Wind)
func (p *RoomContainer) Remove(roomId int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	//	查找并获取房间
	room, err := p.Get(roomId)

	//	如果房间未找到或者房间对象为空
	//	则返回错误
	if err != nil {
		return err
	}

	//	否则清理容器中的房间
	delete(p.rooms, roomId)

	//	将房间内的所有连接都关闭
	for _, session := range room.sessions {
		session.GetConnection().Close()
	}

	//	返回成功
	return nil
}

//	通过房间号获取一个房间
//
//	Author(Wind)
func (p *RoomContainer) Get(roomId int) (*Room, error) {
	//	查找并获取房间
	room, ok := p.rooms[roomId]

	//	如果房间未找到或者房间对象为空
	//	则返回错误
	if !ok || room == nil {
		return nil, ErrContainerNotFoundRoom
	}

	//	返回成功
	return room, nil
}
