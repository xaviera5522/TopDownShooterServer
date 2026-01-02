package session

import (
	"sync"

	"github.com/xaviera5522/TopDownShooterServer/user"
)

type Handler struct {
}

func startSession(user user.User) sync.Map {
	session := newSession(&user)

	var session_map sync.Map
	session_map.Store("id", session.id)
	session_map.Store("last_time", session.lastTime.Format("2006-01-02T15:04:05 -07:00:00"))

	session_user := session.user
	session_map.Store("user_id", session_user.id)
}
