package Services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math"
)

func GenerateToken(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip 1 extra character we get from odd length results
}

func GenerateHash(text string) string {

	byte := md5.Sum([]byte(text))
	hash := hex.EncodeToString(byte[:])

	return hash
}
