// @Author Ben.Zheng
// @DateTime 2022/7/18 4:11 PM

package collections

import "unsafe"

// String2Bytes
// The reason here doesn't use the pointer is that
// pointer will make the variable escape to heap then
// GC will suffer from work pressure.
// Pass a new one, it will be created at stack is faster than
// pointer.
func String2Bytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		Data string
		Cap  int
		Len  int
	}{
		Data: src,
		Cap:  len(src),
		Len:  len(src),
	}))
}

func Bytes2String(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}
