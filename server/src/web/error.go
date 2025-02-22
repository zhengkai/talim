package web

import (
	"net/http"
)

func errorNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"not-found"}`))
}

func errorServerFail(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error":"server-fail"}`))
}
