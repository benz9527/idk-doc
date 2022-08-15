// @Author Ben.Zheng
// @DateTime 2022/8/15 8:58

package intf

type IBiz interface {
	BizError() error
}

type BizPrototype[T IBiz] func() T
