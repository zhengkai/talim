package web

import (
	"net/http"
	"project/config"
	"project/zj"
	"time"
)

// Server ...
func Server() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/api/upload`, uploadHandle)

	s := &http.Server{
		Addr:         config.WebAddr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	zj.J(`start web server`, s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		zj.W(err)
		return
	}
}
