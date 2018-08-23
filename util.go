package zaif

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"
)

func GetNonce() string {
	return fmt.Sprintf("%.6f", float64(time.Now().UnixNano())/float64(time.Second))
}

func Sign(params string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(params))
	return hex.EncodeToString(h.Sum(nil))
}