package testhelpers_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thegodwinproject/cli/internal/testhelpers"
)

func Test_EnsureBuffaloCMD(t *testing.T) {
	r := require.New(t)

	binary := "buffalointegrationtests"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	path := filepath.Join(os.TempDir(), binary)
	os.Remove(path)

	r.NoError(testhelpers.EnsureBuffaloCMD(t))
	r.FileExists(path)
}

//
//func Test_InstallOldBuffaloCMD(t *testing.T) {
//	if runtime.GOOS != "linux" {
//		// this test calls api.github.com too much.
//		// no reason for linux, but just reduce the calls as 1/3
//		t.Skip("skipping Test_InstallOldBuffaloCMD (non-linux env)")
//	}
//
//	tt := []struct {
//		name    string
//		version string
//		err     error
//	}{
//		{
//			name:    "non-existing version",
//			version: "v0.16.40",
//			err:     errors.New("unknown gobuffalo cli version v0.16.40"),
//		},
//		{
//			name:    "existing version",
//			version: "v0.18.1",
//			err:     nil,
//		},
//		{
//			name:    "existing old version",
//			version: "v0.16.27",
//			err:     nil,
//		},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			r := require.New(t)
//
//			err := testhelpers.InstallOldBuffaloCMD(t, tc.version)
//			if tc.err != nil {
//				r.EqualError(err, tc.err.Error())
//				return
//			}
//
//			r.NoError(err)
//			cmd := exec.Command("buffalo", "version")
//			out, err := cmd.CombinedOutput()
//			r.NoError(err)
//			r.Contains(string(out), tc.version)
//		})
//	}
//}
