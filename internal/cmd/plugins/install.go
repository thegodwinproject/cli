package plugins

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/meta"
	"github.com/spf13/cobra"
	"github.com/thegodwinproject/cli/internal/genny/plugins/install"
	"github.com/thegodwinproject/cli/internal/plugins/plugdeps"
)

var installOptions = struct {
	dryRun  bool
	vendor  bool
	verbose bool
}{}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs plugins listed in config/buffalo-plugins.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		if installOptions.dryRun {
			run = genny.DryRunner(context.Background())
			if installOptions.vendor {
				run.FileFn = func(f genny.File) (genny.File, error) {
					bb := &bytes.Buffer{}
					if _, err := io.Copy(bb, f); err != nil {
						return f, err
					}
					return genny.NewFile(f.Name(), bb), nil
				}
			}
		}

		app := meta.New(".")
		plugs, err := plugdeps.List(app)
		if err != nil && !errors.Is(err, plugdeps.ErrMissingConfig) {
			return err
		}

		for _, a := range args {
			plugs.Add(plugdeps.NewPlugin(a))
		}
		gg, err := install.New(&install.Options{
			App:     app,
			Plugins: plugs.List(),
			Vendor:  installOptions.vendor,
		})
		if err != nil {
			return err
		}
		run.WithGroup(gg)

		if installOptions.verbose {
			run.Logger = logger.New(logger.DebugLevel)
		}

		return run.Run()
	},
}

func init() {
	installCmd.Flags().BoolVarP(&installOptions.dryRun, "dry-run", "d", false, "dry run")
	installCmd.Flags().BoolVarP(&installOptions.verbose, "verbose", "v", false, "turn on verbose logging")
	installCmd.Flags().BoolVar(&installOptions.vendor, "vendor", false, "will install plugin binaries into ./plugins [WINDOWS not currently supported]")
}
