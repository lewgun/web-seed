package misc

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)



func MD5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//Base64
func Base64(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))
}

//UnBase64
func UnBase64(src string) string {
	data, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(data)
}

//unused just a helper to make compiler happy
func Unused(...interface{}) {}
