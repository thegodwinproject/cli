package version

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thegodwinproject/cli/internal/runtime"
)

func run(c *cobra.Command, args []string) {
	if !jsonOutput {
		logrus.Infof("Buffalo version is: %s", runtime.Version)
		return
	}

	build := runtime.BuildInfo{}
	build.Version = runtime.Version

	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "    ")
	enc.Encode(build)
}
