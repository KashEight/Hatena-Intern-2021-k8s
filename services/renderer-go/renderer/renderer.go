package renderer

import (
	"bytes"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"

	pb "github.com/hatena/Hatena-Intern-2021/services/renderer-go/pb/fetcher"
)

var markdown = goldmark.New(
	goldmark.WithParserOptions(
		parser.WithASTTransformers(
			util.Prioritized(&autoTitleLinker{}, 999),
		),
	),
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
	if err := markdown.Convert([]byte(newSrc), &w); err != nil {
		return "", errors.New("error")
	}

	html := w.String()

	return html, nil
}

type autoTitleLinker struct {
	fetcherCli pb.FetchererClient
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			node.AppendChild(node, ast.NewString([]byte(l.fetchTitle(string(node.Destination)))))
		}
		return ast.WalkContinue, nil
	})
}

func (l *autoTitleLinker) fetchTitle(url string) string {
	reply, err := l.fetcherCli.Fetcher(context.Background(), &pb.FetcherRequest{Url: url})
	if err != nil {
		return "Invalid"
	}
	return reply.Title
}