package config

import (
	"os"
	"path/filepath"
)

func init() {

	Dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	list := map[string]*string{
		`TALIM_WEB`:   &WebAddr,
		`TALIM_MYSQL`: &MySQL,
		`TALIM_DIR`:   &StaticDir,
		`TALIM_TOKEN`: &Token,
	}
	for k, v := range list {
		s := os.Getenv(k)
		if len(s) > 1 {
			*v = s
		}
	}
}
