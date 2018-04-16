package comment

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
