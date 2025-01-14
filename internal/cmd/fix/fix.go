package fix

import (
	"context"
	"os"
	"os/signal"

	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/thegodwinproject/cli/internal/genny/fix"
)

// run all compatible checks
func RunE(cmd *cobra.Command, args []string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := &fix.Options{
		YesToAll: yesToAll,
	}

	run := genny.WetRunner(ctx)
	if err := run.WithNew(fix.New(opts)); err != nil {
		return err
	}
	return run.Run()
}
