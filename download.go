package download

import (
	"io"
	"net/http"
	"os"
)

func Download(src, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	r, err := http.Get(src)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	_, err = io.Copy(out, r.Body)
	if err != nil {
		return err
	}
	return nil
}
