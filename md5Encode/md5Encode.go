package md5Encode

import (
	"crypto/md5"
	"fmt"
)

func Md5Encode(s string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(s))
	md5 := hash.Sum(nil)

	return fmt.Sprintf("%x", md5)
}
