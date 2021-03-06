package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Handler - stores pointer to our comments service
type Handler struct {
	Router *mux.Router
}

//New Handler -returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

//Setup Routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Rotues")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
