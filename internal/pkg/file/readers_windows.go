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

	fp = strings.ReplaceAll(fp, "/", "\\")

	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		panic(err)
	}

	var path, filename, ext string
	if strings.HasPrefix(fp, ".\\") {
		path, _ = filepath.Abs(path) // Missing file path parameter suffix.
	}

	last := strings.LastIndex(fp, "\\")
	path, filename = path+"\\"+fp[2:last], fp[last+1:]
	if strings.Contains(filename, ".") {
		res := strings.Split(filename, ".")
		filename = res[0]
		ext = res[1]
	} else {
		panic(fmt.Errorf("unable to parse file path [%s] with uncompleted info", fp))
	}

	switch strings.ToLower(ext) {
	case "yaml", "yml":
		return newYamlReader(path, filename, ext)
	case "toml":
		return newTomlReader(path, filename, ext)
	default:
		panic(fmt.Errorf("unknown and not support file extension [%s]", ext))
	}
	return nil
}
