package web

import (
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {

	o := newView(r)
	if o == nil {
		errorNotFound(w)
		return
	}

	re := o.Index()
	if re == nil {
		errorNotFound(w)
		return
	}

	writeJSON(w, re)
}
