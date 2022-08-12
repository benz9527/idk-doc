// @Author Ben.Zheng
// @DateTime 2022/8/5 15:07

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
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

	if strings.HasPrefix(dir, ".\\") {
		wd, e := cast.ToStringE(viper.Get(consts.APP_ROOT_WORKING_DIR))
		if e != nil || wd == "" {
			wd, _ = filepath.Abs(".") // Missing file dir parameter suffix.
		}
		dir = filepath.Join(wd, dir[2:])
	} else if strings.HasPrefix(dir, "..\\") { // Get upper dir absolute string.
		var err error
		if _, err = regexp.MatchString(`^(\.\.\\)+`, dir); err != nil {
			panic(err)
		}
		upperLen := strings.LastIndex(dir, "..\\") + 3
		upperDir := dir[:upperLen]
		if upperDir, err = filepath.Abs(upperDir); err != nil {
			panic(err)
		}
		dir = filepath.Join(upperDir, dir[upperLen:last])
	}

	if _, e := regexp.MatchString(`^[A-Za-z]:\\`, dir); e != nil {
		panic(e)
	}

	finalPath := filepath.Join(dir, filenameWithExt)
	_, err := os.Stat(finalPath)
	if os.IsNotExist(err) {
		panic(err)
	}

	switch strings.ToLower(ext) {
	case "toml":
		return newTomlReader(viper, dir, filename, ext)
	case "yaml", "yml":
		return newYamlReader(viper, dir, filename, ext)
	default:
		panic(fmt.Errorf("unknown and not support file extension [%s]", ext))
	}
	return nil
}
