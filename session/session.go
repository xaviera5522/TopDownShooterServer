package session

import (
	"time"

	"github.com/google/uuid"
	"github.com/xaviera5522/TopDownShooterServer/user"
)

const INACTIVE_TIMEOUT_MINUTES = 3.0

type Session struct {
	id       string
	user     *user.User
	lastTime time.Time
}

func newSession(user *user.User) *Session {
	return &Session{
		user:     user,
		id:       uuid.NewString(),
		lastTime: time.Now(),
	}
}

func (s *Session) IsInactive() bool {
	return time.Since(s.lastTime).Minutes() >= INACTIVE_TIMEOUT_MINUTES
}
