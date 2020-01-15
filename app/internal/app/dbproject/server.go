package dbproject

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ImbaCow/bd_project/internal/app/model"
	socketio "github.com/googollee/go-socket.io"

	"github.com/ImbaCow/bd_project/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName = "auth"
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
	socketServer *socketio.Server
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()
	s.configureSocketServer()
	return s
}

//Start ...
func (s *server) Start(config *Config) error {
	if err := s.configureLogger(config.LogLevel); err != nil {
		return err
	}
	s.logger.Info("Starting api server")

	go s.socketServer.Serve()
	defer s.socketServer.Close()

	err := http.ListenAndServe(config.BindAddr, s.router)
	if err != nil {
		s.logger.Error(err)
	}
	return err
}

func (s *server) configureLogger(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/user/add", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")
	s.router.HandleFunc("/channels/all", s.handleGetChanels())
}

func (s *server) configureSocketServer() {
	s.socketServer = newSocketHandler(s)
	s.router.Handle("/socket", s.socketServer)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Login:    req.Login,
			Password: req.Password,
		}
		id, err := s.store.GetRepositoryStorage().GetUserRepository().Create(u)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		} else {
			s.respond(w, r, http.StatusCreated, map[string]int{"id": id})
		}
	}
}

func (s *server) handleLogin() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	incorectLoginOrPasswordError := errors.New("Incorrect login or password")
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.GetRepositoryStorage().GetUserRepository().FindByLogin(req.Login)
		if err != nil || !u.IsPasswordEqual(req.Password) {
			s.error(w, r, http.StatusUnauthorized, incorectLoginOrPasswordError)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		session.Values["id"] = u.ID
		session.Values["login"] = u.Login
		if err := session.Save(r, w); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleGetChanels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channelRepo := s.store.GetRepositoryStorage().GetChannelRepository()
		channels, err := channelRepo.FindAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusOK, channels)
	}
}
