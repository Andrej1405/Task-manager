package app

import (
	"app/app/model/entities"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var Cache = new(cache)

// Структура кэша для хранения сессий пользователя
type cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	sessions          map[string]entities.Session
}

func (c *cache) init() {
	c.cleanupInterval = 1500 * time.Second
	c.defaultExpiration = 1700 * time.Second
	c.sessions = make(map[string]entities.Session)
}

// Возвращает true, если такая сессия существует в кэше
func GetSessionById(id string) bool {
	var exist bool
	_, exist = Cache.sessions[id]

	return exist
}

// Добавление новой сесии в кэш
func Add(sid string, email string) error {
	if Cache.cleanupInterval == 0 || Cache.defaultExpiration == 0 || Cache.sessions == nil {
		Cache.init()
	}

	var created = time.Now()
	var duration = Cache.defaultExpiration
	var expiration = time.Now().Add(duration).UnixNano()

	Cache.Lock()
	defer Cache.Unlock()

	Cache.sessions[sid] = entities.Session{
		Email:      email,
		Created:    created,
		Expiration: expiration,
	}

	return nil
}

func DeleteBySID(sid string) error {
	Cache.Lock()
	defer Cache.Unlock()

	delete(Cache.sessions, sid)

	return nil
}
