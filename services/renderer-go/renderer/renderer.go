package renderer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
)

// var urlRE = regexp.MustCompile(`https?://[^\s]+`)
// var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))

var counter = 0

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	// Increment counter and replace

	counter++

	if strings.Contains(src, "<+>") {
		_ = strings.Replace(src, "<+>", fmt.Sprintf("%07d", counter), -1)
	}

	// Convert Markdown

	var w bytes.Buffer
	if err := goldmark.Convert([]byte(src), &w); err != nil {
		return "", errors.New("error")
	}

	html := w.String()

	return html, nil
}
