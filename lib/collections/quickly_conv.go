// @Author Ben.Zheng
// @DateTime 2022/7/18 4:11 PM

package collections

import "unsafe"

func String2Bytes(src *string, len int) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		Data string
		Cap  int
		Len  int
	}{
		Data: *src,
		Cap:  len,
		Len:  len,
	}))
}

func Bytes2String(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}
