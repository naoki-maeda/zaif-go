package zaif

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func DecodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func GetNonce() string {
	return fmt.Sprintf("%.6f", float64(time.Now().UnixNano())/float64(time.Second))
}

func Sign(params string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(params))
	return hex.EncodeToString(h.Sum(nil))
}
