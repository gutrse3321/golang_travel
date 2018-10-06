package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

//session接口
//实现session的设置值、读取值、删除值、获取当前sessionID
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// session管理器底层储存结构
type Provider interface {
	SessionInit(sid string) (Session, error) // 实现Session的初始化，操作成功者返回此新的Session变量
	SessionRead(sid string) (Session, error) // 返回sid所代表的Session变量，如果不存在，那么将以sid为参数调用SessionInit函数，创建新的SSession
	SessionDestroy(sid string) error         // 销毁sid对应的Session变量
	SessionGC(maxLifeTime int64)             // 根据maxLifeTime来删除过期的数据
}

// 定义一个全局的session管理器
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

//先定义好接口，然后具体的存储 session 的结构实现相应的接口并注册后，
//相应功能这样就可以使用了，以下是用来随需注册存储 session 的结构的 Register 函数的实现。
var provides = make(map[string]Provider)

func Register(name string, provide Provider) {
	if provide == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provide
}

func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

// 全局唯一的SessionID
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	// 返回base64转换的字符串
	return base64.URLEncoding.EncodeToString(b)
}

// session创建
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// 获取cookie值
	cookie, err := r.Cookie(manager.cookieName)
	// 判断是否有cookie
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
}
