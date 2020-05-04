package app

import (
	"app/app/model/entities"
	"sync"

	_ "github.com/lib/pq"
)

var Cache = new(cache)

// Структура кэша для хранения сессий пользователя
type cache struct {
	sync.RWMutex
	sessions map[string]entities.Session
}

func (c *cache) init() {
	c.sessions = make(map[string]entities.Session)
}

// Возвращает true, если такая сессия существует в кэше
func GetSessionById(id string) bool {
	var boolean bool
	_, boolean = Cache.sessions[id]

	return boolean
}

// Добавление новой сесии в кэш
func Add(sid string, email string) error {
	if Cache.sessions == nil {
		Cache.init()
	}

	Cache.Lock()
	defer Cache.Unlock()

	Cache.sessions[sid] = entities.Session{
		Email: email,
	}

	return nil
}

func DeleteBySID(sid string) error {
	Cache.Lock()
	defer Cache.Unlock()

	delete(Cache.sessions, sid)

	return nil
}
