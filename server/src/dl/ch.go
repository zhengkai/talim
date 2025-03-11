package dl

import "project/config"

var ch chan string

func Insert(s string) {

	if !config.Prod {
		return
	}

	if s == `` {
		return
	}

	select {
	case ch <- s:
	default:
	}

}

func Loop() {
	ch = make(chan string, 10000)
	for v := range ch {
		download(v)
	}
}
