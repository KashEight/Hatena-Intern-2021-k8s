package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render(t *testing.T) {
	src :=
		`# hoge

	- foo
	- bar

	[url](https://example.com)`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t,
		`<h1>hoge</h1>\n<ul>\n<li>foo</li>\n<li>bar</li>\n</ul>\n<p><a href=\"https://example.com\">url</a></p>`,
		html)
}
