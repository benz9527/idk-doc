// @Author Ben.Zheng
// @DateTime 2022/8/5 15:46

package file

import (
	"fmt"
	"log"

	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Only support .yaml and .yml configuration file.
// If there are others file type, we can implement the read details
// then set the value into viper.

type yamlReader struct {
	viper *viper.Viper
}

func newYamlReader(viper *viper.Viper, dir, filename, extension string) intf.IConfigurationReader {
	v := viper
	// Config dir for viper, it will make viper read all files
	// under the specific dir. We should point out which file
	// should be read.
	v.AddConfigPath(dir)
	v.SetConfigType(extension)
	// Windows system uses the backslash, but below are also okay in viper.
	v.SetConfigFile(fmt.Sprintf("%s/%s.%s", dir, filename, extension))
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("read from file [%s.%s] under dir [%s] with error [%v]", filename, extension, dir, err)
	}
	return &yamlReader{
		viper: v,
	}
}

func (y *yamlReader) GetAny(key string) (value any, err error) {
	return y.viper.Get(key), nil
}

func (y *yamlReader) GetString(key string) (value string, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToStringE(a)
}

func (y *yamlReader) GetBoolean(key string) (value bool, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToBoolE(a)
}

func (y *yamlReader) GetInt64(key string) (value int64, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToInt64E(a)
}

func (y *yamlReader) GetFloat64(key string) (value float64, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToFloat64E(a)
}

func (y *yamlReader) GetSlice(key string) (value []any, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToSliceE(a)
}

func (y *yamlReader) GetMap(key string) (value map[string]any, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	// Yaml integer numeric will be parsed into 64/32bit (int)
	return cast.ToStringMapE(a)
}
