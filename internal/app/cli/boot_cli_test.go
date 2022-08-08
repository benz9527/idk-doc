// @Author Ben.Zheng
// @DateTime 2022/8/8 15:28

package cli

import (
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
)

func Test_app_boot_cli_with_help(t *testing.T) {
	asserter := assert.New(t)

	args := []string{
		"--help",
	}
	opts := &BootOptions{}
	if err := opts.Parse(args); err != nil {
		if e, ok := err.(*flags.Error); !ok {
			asserter.NotNil(e)
		} else {
			asserter.Equal(flags.ErrHelp, e.Type)
		}
	}
}

func Test_app_boot_cli_with_cfg_file_path_short_opt(t *testing.T) {
	asserter := assert.New(t)

	expectedPath := "./"
	args := []string{
		"-f=" + expectedPath,
	}
	opts := &BootOptions{}
	if err := opts.Parse(args); err != nil {
		if e, ok := err.(*flags.Error); !ok {
			asserter.NotNil(e)
		} else {
			asserter.NotEqual(flags.ErrHelp, e.Type)
		}
	}
	asserter.Equal(expectedPath, opts.FilePath)
}
