package plugins

import (
	"context"
	"errors"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/meta"
	"github.com/spf13/cobra"
	"github.com/thegodwinproject/cli/internal/genny/add"
	"github.com/thegodwinproject/cli/internal/plugins/plugdeps"
)

var addOptions = struct {
	dryRun    bool
	buildTags []string
}{}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds plugins to config/buffalo-plugins.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		if addOptions.dryRun {
			run = genny.DryRunner(context.Background())
		}

		app := meta.New(".")
		plugs, err := plugdeps.List(app)
		if err != nil && !errors.Is(err, plugdeps.ErrMissingConfig) {
			return err
		}

		tags := app.BuildTags("", addOptions.buildTags...)
		for _, a := range args {
			plugs.Add(plugdeps.NewPlugin(a, tags))
		}
		g, err := add.New(&add.Options{
			App:     app,
			Plugins: plugs.List(),
		})
		if err != nil {
			return err
		}
		if err := run.With(g); err != nil {
			return err
		}

		return run.Run()
	},
}

func init() {
	addCmd.Flags().BoolVarP(&addOptions.dryRun, "dry-run", "d", false, "dry run")
	addCmd.Flags().StringSliceVarP(&addOptions.buildTags, "tags", "t", []string{}, "build tags")
}
