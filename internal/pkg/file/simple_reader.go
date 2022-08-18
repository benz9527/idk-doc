// @Author Ben.Zheng
// @DateTime 2022/8/11 13:22

package file

// For idk internal tests.

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

type simpleReader struct {
	viper *viper.Viper
}

func NewSimpleReader(v *viper.Viper) intf.IConfigurationReader {
	return &simpleReader{
		viper: v,
	}
}
func (y *simpleReader) GetAny(key string) (value any, err error) {
	return y.viper.Get(key), nil
}

func (y *simpleReader) GetString(key string) (value string, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToStringE(a)
}

func (y *simpleReader) GetBoolean(key string) (value bool, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToBoolE(a)
}

func (y *simpleReader) GetInt64(key string) (value int64, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToInt64E(a)
}

func (y *simpleReader) GetFloat64(key string) (value float64, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToFloat64E(a)
}

func (y *simpleReader) GetSlice(key string) (value []any, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	return cast.ToSliceE(a)
}

func (y *simpleReader) GetMap(key string) (value map[string]any, err error) {
	a, err := y.GetAny(key)
	if err != nil {
		return value, err
	}
	// Yaml integer numeric will be parsed into 64/32bit (int)
	return cast.ToStringMapE(a)
}
