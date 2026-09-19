package hello

import (
	"fmt"
	"net/http"
	"os"

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

func (h *handler) Env(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Deployd will give you these additional ENV variables your appplication can use:\n\n")
	fmt.Fprintf(w, "DEPLOYD_SECRET=%s\n", os.Getenv("DEPLOYD_SECRET"))
	fmt.Fprintf(w, "DEPLOYD_RAFT=%s\n", os.Getenv("DEPLOYD_RAFT"))
	fmt.Fprintf(w, "DEPLOYD_SERVICE_NAMESPACE=%s\n", os.Getenv("DEPLOYD_SERVICE_NAMESPACE"))
	fmt.Fprintf(w, "DEPLOYD_SERVICE=%s\n", os.Getenv("DEPLOYD_SERVICE"))
	fmt.Fprintf(w, "DEPLOYD_HOST=%s\n", os.Getenv("DEPLOYD_HOST"))
	fmt.Fprintf(w, "DEPLOYD_HOST_REPLICA_ID=%s\n", os.Getenv("DEPLOYD_HOST_REPLICA_ID"))
	fmt.Fprintf(w, "DEPLOYD_SERVICE_BUILD_COMMIT_ID=%s\n", os.Getenv("DEPLOYD_SERVICE_BUILD_COMMIT_ID"))
	fmt.Fprintf(w, "DEPLOYD_SERVICE_BUILD_ID=%s\n", os.Getenv("DEPLOYD_SERVICE_BUILD_ID"))
	fmt.Fprintf(w, "DEPLOYD_SERVICE_BUILD_DATE=%s\n", os.Getenv("DEPLOYD_SERVICE_BUILD_DATE"))
	fmt.Fprintf(w, "DEPLOYD_SERVICE_BUILD_TAG=%s\n", os.Getenv("DEPLOYD_SERVICE_BUILD_TAG"))
}
