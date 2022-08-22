// @Author Ben.Zheng
// @DateTime 2022/8/22 13:11

package db

import (
	crand "crypto/rand"
	"errors"
	"sync"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

var (
	ErrInvalidLen = errors.New("len for ID generator is out of index (2-255)")
	std           = [...]byte{
		'A', 'B', 'C', 'D', 'E',
		'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y',
		'Z', 'a', 'b', 'c', 'd',
		'e', 'f', 'g', 'h', 'i',
		'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x',
		'y', 'z', '0', '1', '2',
		'3', '4', '5', '6', '7',
		'8', '9', '-', '_',
	}
)

func StandardNanoId(len int) (intf.IDGen, error) {
	if len < 2 || len > 255 {
		return nil, ErrInvalidLen
	}

	var size, offset int
	var b, id []byte
	size = len * len * 7
	b = make([]byte, size)
	if _, err := crand.Read(b); err != nil {
		return nil, ErrInvalidLen
	}
	offset, id = 0, make([]byte, len)

	var lock = sync.Mutex{}
	return func() string {
		lock.Lock()
		defer lock.Unlock()

		if offset == size {
			if _, err := crand.Read(b); err != nil {
				return ""
			}
			offset = 0
		}

		for i := 0; i < len; i++ {
			id[i] = std[b[i+offset]&63]
		}
		offset += len
		return string(id)
	}, nil
}
