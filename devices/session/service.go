package session

import (
	"errors"
	"sync"
)

type Service struct {
	//	Session 存储结构
	//	Map<string>Session
	sessions *sync.Map
}

var (
	service *Service
	once    sync.Once
)

//	获取 session service 实例
//
//	Author(Wind)
func GetService() *Service {
	return service
}

//	创建一个存储 Session 服务的对象
//
//	Author(Wind)
func NewService() *Service {
	once.Do(func() {
		service = &Service{}
	})

	return service
}

//	设置 Session
//
//	Author(Wind)
func (p *Service) Set(s *Session) error {
	//	存储 Session
	p.sessions.Store(s.GetSessionId(), *s)

	//	返回成功
	return nil
}

//	删除一个 Session
//
//	Author(Wind)
func (p *Service) Remove(s *Session) error {
	//	关闭 wss 连接
	_ = s.conn.Close()

	//	删除 Session
	p.sessions.Delete(s.GetSessionId())

	//	返回成功
	return nil
}

//	查找一个会话对象
//
//	Author(Wind)
func (p *Service) Find(sessionId string) (*Session, error) {
	//	查找会话对象并断言成会话对象
	if sess, ok := p.sessions.Load(sessionId); ok {
		if s, ok := sess.(Session); ok {
			return &s, nil
		}
	}

	//	未找到会话
	return nil, errors.New("未找到该会话")
}

//	获取所有的 Session 对象
//	排除某一个指定的 Session
//
//	Author(Wind)
func (p *Service) GetAll(except *Session) (result []Session) {
	//	循环整个 Map 整理会话数据
	p.sessions.Range(func(key, value interface{}) bool {
		//	排除自己
		if key.(string) == except.sessionId {
			return true
		}

		//	追加返回值
		result = append(result, value.(Session))

		//	返回成功
		return true
	})

	//	返回成功
	return
}

//	创建服务
func (p *Service) Start() error {
	//	构建新的存储器
	p.sessions = &sync.Map{}

	//	返回成功
	return nil
}

//	结束
func (p *Service) End() error {
	//	TODO: 持久化过程
	return nil
}
