package api

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/thegodwinproject/cli/internal/genny/newapp/core"
)

//go:embed templates/*
var templates embed.FS

// Templates used for api project files
// (exported mostly for the "fix" command)
func Templates() (fs.FS, error) {
	return fs.Sub(templates, "templates")
}

// New generator for creating a Buffalo API application
func New(opts *Options) (*genny.Group, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	gg, err := core.New(opts.Options)
	if err != nil {
		return gg, err
	}

	g := genny.New()
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
	return gg, nil
}
