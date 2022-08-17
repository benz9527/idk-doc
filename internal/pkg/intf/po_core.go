// @Author Ben.Zheng
// @DateTime 2022/8/17 10:44

package intf

type ICore[T any] interface {
	GetCore() T
}
