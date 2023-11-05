package handler

import (
	"fmt"
	"net/http"
)

func Message(w http.ResponseWriter, r *http.Request) {
	jsonStr := []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	fmt.Print(w, jsonStr)
}
