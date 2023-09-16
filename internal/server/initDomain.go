package server

import (
	memoHandler "memo/internal/domain/memo/handler"
	memoRepository "memo/internal/domain/memo/repository"
	memoUsecase "memo/internal/domain/memo/usecase"
	userHandler "memo/internal/domain/user/handler"
	userRepository "memo/internal/domain/user/repository"
	userUsecase "memo/internal/domain/user/usecase"
	"memo/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func (s *Server) InitDomains() {
	s.initMemo()
}

func (s *Server) initMemo() {
	memoRepository := memoRepository.New(s.ent)
	memoUseCase := memoUsecase.New(
		memoRepository,
	)
	memoHandler := memoHandler.New(memoUseCase)
	s.router.Route("/api/v1/memo", func(router chi.Router) {
		router.Use(middleware.Authenticate)
		router.Post("/create", memoHandler.Create)
		router.Get("/", memoHandler.List)
	})

	userRepository := userRepository.New(s.ent)
	userUseCase := userUsecase.New(
		userRepository,
	)
	userHandler := userHandler.New(userUseCase)
	s.router.Route("/api/v1/user", func(router chi.Router) {
		router.Post("/register", userHandler.Register)
		router.Post("/login", userHandler.Login)
		router.Get("/logout", userHandler.Logout)
		router.Get("/list", userHandler.List)
	})
}
