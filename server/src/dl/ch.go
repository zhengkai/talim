package dl

import "project/config"

var ch = make(chan string)

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
	for v := range ch {
		download(v)
	}
}
