package security

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))

	sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return fmt.Sprint(sha)
}
