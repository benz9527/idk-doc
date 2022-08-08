// @Author Ben.Zheng
// @DateTime 2022/8/5 15:07

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

func NewConfigurationReader(fp string) intf.IConfigurationReader {

	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		panic(err)
	}

	if strings.Contains(fp, "\\") {
		panic(fmt.Errorf("unknown backslash for config file path [%s] in linux", fp))
	}

	var dir, filename, ext string
	last := strings.LastIndex(fp, "/")
	if strings.HasPrefix(fp, "./") {
		dir, _ = filepath.Abs(dir)
		dir = dir + "/" + fp[2:last]
	}

	filename = fp[last+1:]
	if strings.Contains(filename, ".") {
		res := strings.Split(filename, ".")
		filename = res[0]
		ext = res[1]
	} else {
		panic(fmt.Errorf("unable to parse file dir [%s] with uncompleted info", fp))
	}

	switch strings.ToLower(ext) {
	case "yaml", "yml":
		return newYamlReader(dir, filename, ext)
	case "toml":
		return newTomlReader(dir, filename, ext)
	default:
		panic(fmt.Errorf("unknown and not support file extension [%s]", ext))
	}
	return nil
}
