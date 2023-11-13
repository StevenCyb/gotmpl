package input

import (
	"errors"
	"net/http"
)

var (
	defaultHTTPClient  httpClientI = http.DefaultClient
	ErrHTTPStatusNotOK             = errors.New("http status not ok")
)

type httpClientI interface {
	Get(url string) (resp *http.Response, err error)
}

func FromWeb(url string) ([]byte, error) {
	resp, err := defaultHTTPClient.Get(url)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, ErrHTTPStatusNotOK
	}

	return FromReader(resp.Body)
}
