package dbproject

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type connectFunc func(socketio.Conn) error

func newSocketHandler(s *server) *socketio.Server {
	server, _ := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	return server
}
