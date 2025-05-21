package helpers

import (
	"fmt"
	"io"
	"net/http"
)


var (
	ErroBadRequest = fmt.Errorf("bad request")
	ErroNotFound   = fmt.Errorf("not found")
	ErroInternal   = fmt.Errorf("internal server error")
	ErroTimeout    = fmt.Errorf("timeout")
	ErroUnknown    = fmt.Errorf("unknown error")
	ErrReadBody   = fmt.Errorf("error reading response body")
)


func FetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody
	}

	return body, nil
}
