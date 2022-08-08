// @Author Ben.Zheng
// @DateTime 2022/8/8 8:46

package file

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/benz9527/idk-doc/internal/pkg/intf"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Only support .toml configuration file.

type tomlReader struct {
	viper *viper.Viper
}

func newTomlReader(dir, filename, extension string) intf.IConfigurationReader {
	v := viper.New()
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
	return &tomlReader{
		viper: v,
	}
}

func (t *tomlReader) GetAny(key string) (value any, err error) {
	hasDot := strings.Contains(key, ".")
	var keys = make([]string, 0, 1)
	if hasDot {
		keys = strings.Split(key, ".")
	} else {
		keys = append(keys, key)
	}

	var result any
	var m map[string]any
	kLen := len(keys)
	for i, k := range keys {
		if i == 0 {
			result = t.viper.Get(k)
		} else {
			var ok bool
			result, ok = m[k]
			if !ok {
				result = nil
			}
		}
		if result == nil {
			return result, errors.New(fmt.Sprintf("Not found key [%s] in toml", key))
		}

		switch result.(type) {
		case []interface{}:
			arr := result.([]interface{})[0]
			switch arr.(type) {
			case map[string]any:
				m = arr.(map[string]any)
				continue
			}
		default:
			if i != kLen-1 {
				return nil, errors.New(fmt.Sprintf("Not found key [%s] in toml", key))
			}
		}
	}
	return result, nil
}

func (t *tomlReader) GetString(key string) (value string, err error) {
	a, err := t.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToStringE(a)
}

func (t *tomlReader) GetBoolean(key string) (value bool, err error) {
	a, err := t.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToBoolE(a)
}

func (t *tomlReader) GetInt64(key string) (value int64, err error) {
	a, err := t.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToInt64E(a)
}

func (t *tomlReader) GetFloat64(key string) (value float64, err error) {
	a, err := t.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToFloat64E(a)
}

func (t *tomlReader) GetSlice(key string) (value []any, err error) {
	a, err := t.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToSliceE(a)
}

func (t *tomlReader) GetMap(key string) (value map[string]any, err error) {
	a, err := t.GetAny(key)
	if err != nil {
		return value, err
	}
	// Toml integer numeric will be parsed into 64 bit (int64)
	return cast.ToStringMapE(a)
}
