package util

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zhengkai/zu"
)

const UserAgent = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36`

func Download(url, file string) (err error) {

	if FileExists(file) {
		return nil
	}

	req, err := http.NewRequest(`GET`, url, nil)
	if err != nil {
		return
	}
	req.Header.Set(`User-Agent`, UserAgent)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf(`bad response status: %s in url %s`, rsp.Status, url)
	}

	f, err := os.CreateTemp(Static(`tmp`), zu.TempFilePattern)
	if err != nil {
		return
	}

	if _, err := io.Copy(f, rsp.Body); err != nil {
		f.Close()
		os.Remove(f.Name())
		return fmt.Errorf(`copy url %s failed: %w`, url, err)
	}

	f.Close()
	Mkdir(file)
	file = Static(file)
	os.Rename(f.Name(), file)
	go os.Chmod(file, 0664)
	return
}
