package fetcher

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetcher(ctx context.Context, url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", errors.New("error")
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", errors.New("error")
	}

	strBody := string(b)
	titleStartIndex := strings.Index(strBody, "<title>")
	titleEndIndex := strings.Index(strBody, "</title>")
	title := strBody[titleStartIndex:titleEndIndex]

	return title, nil
}