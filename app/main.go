package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"

	"github.com/kelseyhightower/envconfig"

	"github.com/ImbaCow/bd_project/internal/app/dbproject"
)

func maintest() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		fmt.Println("connected:", url.Query().Get("channel"))
		for name, values := range s.RemoteHeader() {
			for _, value := range values {
				fmt.Println(name, value)
			}
		}
		s.SetContext("")
		return nil
	})
	server.OnEvent("/", "msg", func(s socketio.Conn) string {
		last := s.Context().(string)
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./src")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	var config dbproject.Config
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(dbproject.Start(&config))
}
