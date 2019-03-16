package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(str string) string {
	encryptInst := sha1.New()
	encryptInst.Write([]byte(str))
	res := hex.EncodeToString(encryptInst.Sum(nil))
	return res
}
