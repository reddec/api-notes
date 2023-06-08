package render

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"

	embed "github.com/13rac1/goldmark-embed"
	"github.com/Masterminds/sprig"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
)

//go:embed marx.min.css
var style string

//go:embed template.gohtml
var templateContent string

func New() *Renderer {
	t := template.Must(template.New("").Funcs(sprig.HtmlFuncMap()).Parse(templateContent))
	return &Renderer{
		page: t,
		engine: goldmark.New(goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
			highlighting.Highlighting,
			&mermaid.Extender{},
			mathjax.MathJax,
			embed.New(),
		),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			)),
	}
}

type Renderer struct {
	page   *template.Template
	engine goldmark.Markdown
}

func (r *Renderer) Render(title, content string, author string, attachments []string) ([]byte, error) {
	var out bytes.Buffer
	if err := r.engine.Convert([]byte(content), &out); err != nil {
		return nil, fmt.Errorf("render: %w", err)
	}
	vd := &viewData{
		Title:       title,
		Author:      author,
		Style:       template.CSS(style),
		HTML:        template.HTML(out.String()),
		Attachments: attachments,
	}
	var page bytes.Buffer
	err := r.page.Execute(&page, vd)
	if err != nil {
		return nil, fmt.Errorf("render HTML: %w", err)
	}
	return page.Bytes(), nil
}

type viewData struct {
	Title       string
	Author      string
	Attachments []string
	Style       template.CSS
	HTML        template.HTML
}
