package game

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/xaviera5522/TopDownShooterServer/session"
)

const MIN_FRAME_TIME_MS = 8.333 //CAP at 120 FPS

type GameServer struct {
	id       string
	sessions []*session.Session
}

func newServer(sessions []*session.Session) *GameServer {
	return &GameServer{
		id:       uuid.NewString(),
		sessions: sessions,
	}
}

func (g *GameServer) GameLLogic() {
	frame_start := time.Now()

	//update user or something

	delta_time := time.Since(frame_start)
	if float64(delta_time.Milliseconds()) < MIN_FRAME_TIME_MS {
		//sleep for the difference
		sleep_duration := MIN_FRAME_TIME_MS - float64(delta_time.Milliseconds())
		time.Sleep(time.Duration(sleep_duration))
		delta_time = time.Since(frame_start)
	}
	fmt.Println("DEBUG: ", delta_time)
}
