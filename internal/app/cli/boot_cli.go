// @Author Ben.Zheng
// @DateTime 2022/8/8 15:02

package cli

import (
	"github.com/jessevdk/go-flags"
)

type BootOptions struct {
	ListenAddr string `short:"A" long:"listen-addr" required:"true" default:"0.0.0.0" description:"Assign listening IP address for idk-doc application to boot up."`
	ListenPort string `short:"p" long:"listen-port" required:"true" default:"9527" description:"Assign listening port number for idk-doc application to boot up."`
	FilePath   string `short:"f" long:"cfg-file-path" required:"true" default:"./conf/idk-boot.yaml" description:"Assign configuration file path for idk-doc application to boot up."`
}

func NewBootOptions() *BootOptions {
	return &BootOptions{}
}

func (o *BootOptions) Parse(args []string) error {
	if _, err := flags.ParseArgs(o, args); err != nil {
		return err
	}
	return nil
}
