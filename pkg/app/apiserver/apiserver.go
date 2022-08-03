package apiserver

import (
	"io"
	"net/http"

	"github.com/gaponovalexey/go-restapi/pkg/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//start
func (s *APIServer) Start() error {
	if err := s.configurateLoger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.Configursore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//loger
func (s *APIServer) configurateLoger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.handleHello())
}
func (s *APIServer) Configursore() error {
	// st := store.New(s.config)
	// if err := st.Open(); err != nil {
		// return err
	// }

	// s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
