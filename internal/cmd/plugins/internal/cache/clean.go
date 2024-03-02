package cache

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thegodwinproject/cli/internal/plugins"
)

// CleanCmd cleans the plugins cache
var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "cleans the plugins cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		os.RemoveAll(plugins.CachePath)
		return nil
	},
}
