package server

import (
	"context"
	"log"
	"net/http"
	"memo/config"
	"memo/ent"

	"github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	cfg    *config.Config
	ent    *ent.Client
	router *chi.Mux
}

type Options func(opts *Server) error

func New() *Server {
	cfg := config.New()
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Init() {
	s.newRouter()
	s.newEnt()
	s.InitDomains()
}

func (s *Server) Run() {
	http.ListenAndServe(":"+s.cfg.Api.Port, s.router)
}

func (s *Server) newRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "http://ui:3031"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	s.router = r
}

func (s *Server) newEnt() {
	entOptions := []ent.Option{}

	// 発行されるSQLをロギング
	// entOptions = append(entOptions, ent.Debug())
	mc := mysql.Config{
		User:                 s.cfg.Database.User,
		Passwd:               s.cfg.Database.Password,
		Net:                  "tcp",
		Addr:                 s.cfg.Database.Host + ":" + s.cfg.Database.Port,
		DBName:               s.cfg.Database.Name,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	client, _ := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	// defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	s.ent = client
}
