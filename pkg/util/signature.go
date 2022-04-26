package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func GetSign(apikey string, seckey string, nonce, timestamp string, body []byte) (r string) {
	key := []byte(seckey)
	mac := hmac.New(sha1.New, key)

	mac.Write([]byte(apikey))
	mac.Write([]byte(nonce))
	mac.Write([]byte(timestamp))
	mac.Write(body)
	r = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}
