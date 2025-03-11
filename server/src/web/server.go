package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/config"
	"project/geoip"
	"project/view"
	"project/zj"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

// Server ...
func Server() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/api/index`, indexHandle)
	mux.HandleFunc(`/api/tweet`, tweetHandle)
	mux.HandleFunc(`/api/recent`, recentHandle)
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

func outputData(w http.ResponseWriter, r *http.Request, v proto.Message) {
	var ab []byte
	var err error
	if r.URL.Query().Get(`output`) == `pb` {
		ab, err = proto.Marshal(v)
		if err != nil {
			errorServerFail(w)
			return
		}
		proto := strings.TrimPrefix(fmt.Sprintf(`%T`, v), `*`)
		mime := fmt.Sprintf(`application/x-protobuf; messageType="%s"`, proto)
		w.Header().Set(`Content-Type`, mime)
	} else {
		ab, err = json.MarshalIndent(v, ``, "\t")
		if err != nil {
			errorServerFail(w)
			return
		}
		w.Header().Set(`Content-Type`, `application/json`)
	}
	w.Header().Set(`Content-Length`, strconv.Itoa(len(ab)))
	w.Write(ab)
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

func newView(w http.ResponseWriter, r *http.Request) *view.View {

	if r.Method != http.MethodGet || r.Header.Get(`Referer`) == `` {
		return nil
	}
	ip := r.Header.Get(`X-Real-IP`)
	if geoip.Check(ip) {
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		return nil
	}
	return view.TheView
}
