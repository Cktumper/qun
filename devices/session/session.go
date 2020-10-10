package session

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

type Session struct {
	expiredAt time.Time
	sessionId string
	conn      *websocket.Conn
}

//	构建 Session
//
//	Author(Wind)
func NewSession() *Session {
	return &Session{}
}

//	创建一个 Session 对象
//
//	Author(Wind)
func (p *Session) Create(conn *websocket.Conn) error {
	//	生成随机数种
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d%d", time.Now().UnixNano(), r.Intn(99999))))

	//	生成一个 sessionId
	p.sessionId = hex.EncodeToString(h.Sum(nil))

	//	Session 存活时间
	p.expiredAt = time.Now().Add(time.Hour * 24 * 30)
	p.conn = conn

	//	返回成功
	return nil
}

//	获取 SessionId
//
//	Author(Wind)
func (p *Session) GetSessionId() string {
	return p.sessionId
}

//	获取 Session 中的连接
//
//	Author(Wind)
func (p *Session) GetConnection() *websocket.Conn {
	return p.conn
}
