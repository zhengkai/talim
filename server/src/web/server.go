package web

import (
	"encoding/json"
	"net/http"
	"project/config"
	"project/util"
	"project/view"
	"project/zj"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

// Server ...
func Server() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/api/index`, indexHandle)
	mux.HandleFunc(`/api/tweet`, tweetHandle)
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

func writeJSON(w http.ResponseWriter, v proto.Message) {
	ab, err := json.MarshalIndent(v, ``, "\t")
	if err != nil {
		errorServerFail(w)
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	w.Header().Set(`Content-Length`, strconv.Itoa(len(ab)))
	w.Write(ab)
}

func newView(r *http.Request) *view.View {
	return view.New(getUUID(r))
}

func getUUID(r *http.Request) (u uuid.UUID) {
	u = util.KaiUUID
	s := r.URL.Query().Get(`uuid`)
	if s == `` {
		return
	}
	pu, err := uuid.Parse(s)
	if err == nil {
		u = pu
	}
	return
}
