package app

import (
	_ "github.com/lib/pq"
)

type Session struct {
	data map[string]string
}

var session = new(Session)

func GetSession(id string) bool {
	var boolean bool
	_, boolean = session.data[id]

	return boolean
}

func Add(id string) error {
	session.data = make(map[string]string)
	session.data[id] = id

	return nil
}

func DeleteSession(id string) error {
	delete(session.data, id)

	return nil
}
