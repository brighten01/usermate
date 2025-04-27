package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

type Encrypt struct {
	Value string
}

// sha256加密和base64 编码
func (e *Encrypt) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(e.Value))
}

func (e *Encrypt) EncryptString() string {
	h := sha256.New()
	h.Write([]byte(e.Value))
	bs := h.Sum(nil)
	return string(bs)
}
