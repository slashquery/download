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

	// create a new request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", src, nil)
	req.Header.Set("User-Agent", "slashquery-download")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}
