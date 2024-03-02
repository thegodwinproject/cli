package web

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/thegodwinproject/cli/internal/genny/assets/standard"
	"github.com/thegodwinproject/cli/internal/genny/assets/webpack"
	"github.com/thegodwinproject/cli/internal/genny/newapp/core"
)

//go:embed templates/* templates/templates/_flash.plush.html.tmpl
var templates embed.FS

// Templates used for web project files
// (exported mostly for the "fix" command)
func Templates() (fs.FS, error) {
	return fs.Sub(templates, "templates")
}

// New generator for creating a Buffalo Web application
func New(opts *Options) (*genny.Group, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	gg, err := core.New(opts.Options)
	if err != nil {
		return gg, err
	}

	g := genny.New()
	g.Transformer(genny.Dot())
	data := map[string]interface{}{
		"opts": opts,
	}

	helpers := template.FuncMap{}

	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)
	sub, err := Templates()
	if err != nil {
		return gg, err
	}

	if err := g.FS(sub); err != nil {
		return gg, err
	}

	gg.Add(g)

	if opts.Webpack != nil {
		// add the webpack generator
		g, err = webpack.New(opts.Webpack)
		if err != nil {
			return gg, err
		}
		gg.Add(g)
	}

	if opts.Standard != nil {
		// add the standard generator
		g, err = standard.New(opts.Standard)
		if err != nil {
			return gg, err
		}
		gg.Add(g)
	}

	return gg, nil
}
