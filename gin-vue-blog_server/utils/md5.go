package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(stc []byte) string {
	m := md5.New()
	m.Write(stc)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
