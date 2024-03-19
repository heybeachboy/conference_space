package util

import (
	"crypto/md5"
	"encoding/base64"
)

func MD5ToString(in string) string {
	md := md5.New()
	md.Write([]byte(in))
	return base64.StdEncoding.EncodeToString(md.Sum(nil))
}
