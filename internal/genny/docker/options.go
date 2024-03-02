package docker

import (
	"github.com/gobuffalo/meta"
	"github.com/thegodwinproject/cli/internal/runtime"
)

type Options struct {
	App     meta.App `json:"app"`
	Version string   `json:"version"`
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts.App.IsZero() {
		opts.App = meta.New(".")
	}

	if len(opts.Version) == 0 {
		opts.Version = runtime.Version
	}

	return nil
}
