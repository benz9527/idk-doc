// @Author Ben.Zheng
// @DateTime 2022/8/5 15:19

package intf

type IConfigurationReader interface {
	GetAny(key string) (value any, err error)
	GetString(key string) (value string, err error)
	GetBoolean(key string) (value bool, err error)
	GetInt64(key string) (value int64, err error)
	GetFloat64(key string) (value float64, err error)
	GetSlice(key string) (value []any, err error)
	GetMap(key string) (value map[string]any, err error)
}
