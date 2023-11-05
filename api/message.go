package handler

import (
	"fmt"
	"net/http"
)

func Message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"message\": \"Hello from Go!\"}")
}
