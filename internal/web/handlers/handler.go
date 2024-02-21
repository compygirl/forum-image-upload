package handlers

import (
	service "forum/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	handlerObj := Handler{service: service}
	return &handlerObj
}

func (handler *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	images := http.FileServer(http.Dir("./data/assets/images"))
	mux.Handle("/images/", http.StripPrefix("/images/", images))
	mux.HandleFunc("/", handler.CheckCookieMiddleware(handler.GetMainPage))
	mux.HandleFunc("/registration", handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.RegistrationHandler)))
	mux.HandleFunc("/login", handler.OnlyUnauthMiddleware(handler.CheckCookieMiddleware(handler.LoginHandler)))
	mux.HandleFunc("/logout", handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.LogoutHandler)))
	mux.HandleFunc("/submit-post", handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.CreatePostHandler)))
	mux.HandleFunc("/post/react", handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ReactOnPostHandler)))
	mux.HandleFunc("/comments/", handler.CheckCookieMiddleware(handler.DisplayCommentsHandler))
	mux.HandleFunc("/submit-comment", handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.CreateCommentsHandler)))
	mux.HandleFunc("/comment/react", handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ReactOnCommentHandler)))
	mux.HandleFunc("/filter/", handler.CheckCookieMiddleware(handler.FilterHandler))
	mux.HandleFunc("/auth/google/in", handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GoogleAuthHandler)))
	mux.HandleFunc("/auth/google/callback", handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GoogleCallback)))
	mux.HandleFunc("/auth/github/in", handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GithubAuthHandler)))
	mux.HandleFunc("/auth/github/callback", handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GithubCallback)))
	return mux
}
