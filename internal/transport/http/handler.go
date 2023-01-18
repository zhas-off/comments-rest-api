package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

// Response object
type Response struct {
	Message string `json:"message"`
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service CommentService) *Handler {
	log.Info("setting up our handler")
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	// Sets up our middleware functions
	h.Router.Use(JSONMiddleware)
	// Log every incoming request
	h.Router.Use(LoggingMiddleware)
	// Timeout all requests that take longer than 15 seconds
	h.Router.Use(TimeoutMiddleware)
	// Set up the routes
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}
	// return handler
	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", h.AliveCheck).Methods("GET")
	h.Router.HandleFunc("/ready", h.ReadyCheck).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")

}

func (h *Handler) AliveCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
		panic(err)
	}
}

func (h *Handler) ReadyCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.ReadyCheck(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "I am Ready!"}); err != nil {
		panic(err)
	}
}

// Serve  serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down gracefully")
	return nil
}
