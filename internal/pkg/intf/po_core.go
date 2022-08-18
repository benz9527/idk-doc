// @Author Ben.Zheng
// @DateTime 2022/8/17 10:44

package intf

type ICore[T comparable] interface {
	// GetCore
	// Avoid to let caller access with useless fields, like foreign key definition field.
	GetCore() T
}
