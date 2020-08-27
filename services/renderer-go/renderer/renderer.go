package renderer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"regexp"

	"errors"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var standAloneUrlRE = regexp.MustCompile(`[^(\(.*\)\[)]https?://[^\s]+[^\]]`)
var exactUrlRE = regexp.MustCompile(`https?://[^\s]+`)
var gamingRE = regexp.MustCompile(`\+\+.+\+\+`)
var gamingREInner = regexp.MustCompile(`[^(\+\+)].+[^(\+\+)]`)
var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var htmlBuf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	if err := md.Convert([]byte(src), &htmlBuf); err != nil {
		return src, errors.New("failed to convert markdown")
	}
	fullConverted := gamingRE.ReplaceAllStringFunc(htmlBuf.String(), func(inner string) string {
		return fmt.Sprintf("<span class=\"gaming\">%s</span>", inner[2:len(inner)-2])
	})
	return fullConverted, nil
}
