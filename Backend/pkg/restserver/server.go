package restserver

import (
	"log"
	"net/http"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/application"
	"github.com/Jmanuel2099/EnronCorpEmail/pkg/datasource/zincsearchdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	port = ":8080"
)

type Server struct {
	ServerChi    *chi.Mux
	httpClient   *http.Client
	emailControl *application.EmailControl
}

func (s *Server) setupServerChi() {
	s.ServerChi = chi.NewRouter()
	s.ServerChi.Use(middleware.RequestID)
	s.ServerChi.Use(middleware.Logger)
	s.ServerChi.Use(middleware.Recoverer)
	s.ServerChi.Use(middleware.URLFormat)
	s.ServerChi.Use(render.SetContentType(render.ContentTypeJSON))
	s.SetupEmailRoutes()
}

func (s *Server) setupHttpClient() {
	s.httpClient = &http.Client{}
}

func (s *Server) setupEmailControl() {
	zincSearchDb := zincsearchdb.NewZincSearchClient(s.httpClient)
	emailContorl := application.NewEmailControl(zincSearchDb)
	s.emailControl = emailContorl
}

func NewServer() *Server {
	server := &Server{}
	server.setupHttpClient()
	server.setupEmailControl()
	server.setupServerChi()

	return server
}

func (s Server) RunServer() {
	log.Printf("App email enron is running on: http://localhost%s\n", port)
	err := http.ListenAndServe(port, s.ServerChi)
	if err != nil {
		panic(err)
	}
}
