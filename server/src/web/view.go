package web

import (
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {

	o := newView(w, r)
	if o == nil {
		errorNotFound(w)
		return
	}

	re := o.Index()
	if re == nil {
		errorNotFound(w)
		return
	}

	outputData(w, r, re)
}

func recentHandle(w http.ResponseWriter, r *http.Request) {

	o := newView(w, r)
	if o == nil {
		errorNotFound(w)
		return
	}

	re := o.TweetRecent(getNum(r, `tid`))
	if re == nil {
		errorServerFail(w)
		return
	}

	outputData(w, r, re)
}

func tweetHandle(w http.ResponseWriter, r *http.Request) {

	o := newView(w, r)
	if o == nil {
		errorNotFound(w)
		return
	}

	re := o.TweetList(getNum(r, `uid`), getNum(r, `tid`))
	if re == nil {
		errorNotFound(w)
		return
	}

	outputData(w, r, re)
}
