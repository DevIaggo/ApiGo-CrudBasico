package api

import (
	"apiGo/internal/handler"
	"net/http"
)

func SetupRoutes(h *handler.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /task", h.CreateTask)
	mux.HandleFunc("GET /task", h.GetAllTasks)
	mux.HandleFunc("GET /task/{id}", h.GetTaskById)
	mux.HandleFunc("DELETE /task/{id}", h.DeleteTaskById)
	mux.HandleFunc("PUT /task/{id}", h.UpdateTask)

	return mux
}
