package build

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"time"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/meta"
	"github.com/spf13/cobra"
	"github.com/thegodwinproject/cli/internal/genny/build"
)

var buildOptions = struct {
	*build.Options
	SkipAssets             bool
	SkipBuildDeps          bool
	Debug                  bool
	Tags                   string
	SkipTemplateValidation bool
	DryRun                 bool
	Verbose                bool
	bin                    string
}{
	Options: &build.Options{
		BuildTime: time.Now(),
	},
}

func runE(cmd *cobra.Command, args []string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	buildOptions.App = meta.New(pwd)
	if len(buildOptions.bin) > 0 {
		buildOptions.App.Bin = buildOptions.bin
	}

	buildOptions.Options.WithAssets = !buildOptions.SkipAssets
	buildOptions.Options.WithBuildDeps = !buildOptions.SkipBuildDeps

	run := genny.WetRunner(ctx)
	if buildOptions.DryRun {
		run = genny.DryRunner(ctx)
	}

	if buildOptions.Verbose || buildOptions.Debug {
		lg := logger.New(logger.DebugLevel)
		run.Logger = lg
		buildOptions.BuildFlags = append(buildOptions.BuildFlags, "-v")
	}

	opts := buildOptions.Options
	opts.BuildVersion = buildVersion(opts.BuildTime.Format(time.RFC3339))

	if buildOptions.Tags != "" {
		opts.Tags = append(opts.Tags, buildOptions.Tags)
	}

	if !buildOptions.SkipTemplateValidation {
		opts.TemplateValidators = append(opts.TemplateValidators, build.PlushValidator, build.GoTemplateValidator)
	}

	if cmd.CalledAs() == "install" {
		opts.GoCommand = "install"
	}
	clean := build.Cleanup(opts)
	// defer clean(run)
	defer func() {
		if err := clean(run); err != nil {
			log.Fatalf("build:clean %s", err)
		}
	}()
	if err := run.WithNew(build.New(opts)); err != nil {
		return err
	}
	return run.Run()
}

func buildVersion(version string) string {
	vcs := buildOptions.VCS

	if len(vcs) == 0 {
		return version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key != "vcs.revision" {
				continue
			}

			return setting.Value
		}
	}

	return version
}
