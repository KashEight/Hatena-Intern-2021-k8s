package renderer

import (
	"bytes"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/yuin/goldmark"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	// Increment counter and replace

	var newSrc string

	newUUID, err := uuid.NewRandom()

	if err != nil {
		return "", errors.New("error")
	}

	if strings.Contains(src, "<+>") {
		newSrc = strings.Replace(src, "<+>", newUUID.String(), -1)
	} else {
		newSrc = src
	}

	// Convert Markdown

	var w bytes.Buffer
	if err := goldmark.Convert([]byte(newSrc), &w); err != nil {
		return "", errors.New("error")
	}

	html := w.String()

	return html, nil
}
