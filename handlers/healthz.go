package handlers

import "net/http"

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.writeHeader(http.StatusOK)
}
