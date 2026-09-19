package hello

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) HelloWorld(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!\n")
}
