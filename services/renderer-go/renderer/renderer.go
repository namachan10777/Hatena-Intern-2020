package renderer

import (
	"bytes"
	"context"
	"html/template"
	"regexp"
	pb_title_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/title_fetcher"
)

var urlRE = regexp.MustCompile(`https?://[^\s\"]+`)
var emptyLinkTagRE = regexp.MustCompile(`<a\s*href\s*=\s*"https:?//.+">\s*<\s*/\s*a\s*>`)
var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.Url}}">{{.Title}}</a>`))

type Link struct {
	Url string
	Title string
}

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, cli pb_title_fetcher.TitleFetcherClient, src string) (string, error) {
	// TODO: これはサンプル実装 (URL の自動リンク) です
	html := urlRE.ReplaceAllStringFunc(src, func(url string) string {
		var w bytes.Buffer
		link := Link {
			Url: url,
			Title: "",
		}
		err := linkTmpl.ExecuteTemplate(&w, "link", link)
		if err != nil {
			return url
		}
		return w.String()
	})
	converted := emptyLinkTagRE.ReplaceAllStringFunc(html, func(a string) string {
		// lenth must be at least 1
		url := urlRE.FindStringSubmatch(a)
		link := Link {
			Url: url[0],
			Title: url[0],
		}
		if title, err := cli.Fetch(ctx, &pb_title_fetcher.FetchRequest{ Url: url[0], }); err == nil {
			link.Title = title.Title
		}
		var w bytes.Buffer
		if err := linkTmpl.ExecuteTemplate(&w, "link", link); err != nil {
			println("err")
			return src
		}
		return w.String()
	})
	return converted, nil
}
