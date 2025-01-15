package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) Healthcheck(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`

	js = fmt.Sprintf(js, h.conf.Environment, h.conf.AppVersion)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js))
}
