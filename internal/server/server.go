package server

import (
	"context"
	"forum/cmd/config"
	repository "forum/internal/database"
	database "forum/internal/database/migration"
	"forum/internal/service"
	handlers "forum/internal/web/handlers"
	"log"
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
}

func InitServer(conf *config.Config, ctx context.Context) *Server {
	db, err := database.CreateDb(conf.DbName, conf.DbPath, ctx)
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.NewRepository(db) // stores the db in the repository
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)

	port := os.Getenv("PORT")
	if port == "" {
		port = conf.Address
	}

	ServerObj := Server{
		httpServer: &http.Server{
			Addr:    port,
			Handler: handler.InitRouter(),
		},
	}

	return &ServerObj
}

func (server *Server) Start() error {
	log.Println("starting api server at http://localhost" + server.httpServer.Addr)
	return server.httpServer.ListenAndServe()
}
