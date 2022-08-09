// @Author Ben.Zheng
// @DateTime 2022/8/5 15:07

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func NewConfigurationReader(viper *viper.Viper, fp string) intf.IConfigurationReader {

	fp = strings.ReplaceAll(fp, "/", "\\")
	last := strings.LastIndex(fp, "\\")

	var dir, filename, ext string
	dir = fp[:last]
	filenameWithExt := fp[last+1:] // Separating the filename with file type extension at first.
	if !strings.Contains(filenameWithExt, ".") {
		panic(fmt.Errorf("unable to parse file path [%s] with uncompleted info", fp))
	}
	res := strings.Split(filenameWithExt, ".")
	filename, ext = res[0], res[1]

	if matched, e := regexp.MatchString(`^[A-Za-z]:\\`, dir); !matched && e != nil {
		panic(e)
	} else {
		if strings.HasPrefix(dir, ".\\") {
			wd, e := cast.ToStringE(viper.Get(consts.APP_ROOT_WORKING_DIR))
			if e != nil || wd == "" {
				dir, _ = filepath.Abs(dir) // Missing file dir parameter suffix.
			} else {
				dir = wd
			}
		}
	}

	finalPath := dir + "\\" + filenameWithExt
	_, err := os.Stat(finalPath)
	if os.IsNotExist(err) {
		panic(err)
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
