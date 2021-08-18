package renderer

import (
	"bytes"
	"context"
	"errors"

	"github.com/yuin/goldmark"
)

// var urlRE = regexp.MustCompile(`https?://[^\s]+`)
// var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	// TODO: これはサンプル実装 (URL の自動リンク) です
	// html := urlRE.ReplaceAllStringFunc(src, func(url string) string {
	// 	var w bytes.Buffer
	// 	err := linkTmpl.ExecuteTemplate(&w, "link", url)
	// 	if err != nil {
	// 		return url
	// 	}

	// 	return w.String()
	// })

	// Convert Markdown
	var w bytes.Buffer
	if err := goldmark.Convert([]byte(src), &w); err != nil {
		return "", errors.New("error")
	}

	html := w.String()

	return html, nil
}
