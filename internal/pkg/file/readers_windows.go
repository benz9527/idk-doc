// @Author Ben.Zheng
// @DateTime 2022/8/5 15:07

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
	"github.com/benz9527/idk-doc/internal/pkg/ioc"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func NewConfigurationReader(viper *viper.Viper, fp string) intf.IConfigurationReader {

	fp = strings.ReplaceAll(fp, "/", "\\")

	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		panic(err)
	}

	var dir, filename, ext string
	if strings.HasPrefix(fp, ".\\") {
		wd, e := cast.ToStringE(viper.Get(ioc.APP_ROOT_WORKING_DIR))
		if e != nil || wd == "" {
			dir, _ = filepath.Abs(dir) // Missing file dir parameter suffix.
		} else {
			dir = wd
		}
	}

	last := strings.LastIndex(fp, "\\")
	dir, filename = dir+"\\"+fp[2:last], fp[last+1:]
	if strings.Contains(filename, ".") {
		res := strings.Split(filename, ".")
		filename = res[0]
		ext = res[1]
	} else {
		panic(fmt.Errorf("unable to parse file path [%s] with uncompleted info", fp))
	}

	switch strings.ToLower(ext) {
	case "yaml", "yml":
		return newYamlReader(viper, dir, filename, ext)
	case "toml":
		return newTomlReader(viper, dir, filename, ext)
	default:
		panic(fmt.Errorf("unknown and not support file extension [%s]", ext))
	}
	return nil
}
