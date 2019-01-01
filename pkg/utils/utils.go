package utils

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"math/big"
	"regexp"
	"strconv"
)

func DefaultQueryPage(c *gin.Context) (int, int) {
	return DefaultQueryForInt(c, "page", 1), DefaultQueryForInt(c, "pageSize", 10)

}

// DefaultQueryForInt returns the keyed url query value if it exists
func DefaultQueryForInt(c *gin.Context, key string, defaultValue int) int {
	if value, ok := strconv.Atoi(c.Query(key)); ok == nil {
		return value
	}
	return defaultValue
}

// DefaultQueryForInt64 returns the keyed url query value if it exists
func DefaultQueryForInt64(c *gin.Context, key string, defaultValue int64) int64 {
	if value, ok := strconv.ParseInt(c.Query(key), 10, 64); ok == nil {
		return value
	}
	return defaultValue
}

// ParamFromId returns the keyed url param value if it exists
func ParamFromID(c *gin.Context, key string) (int64, error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandomString returns generated random string in given length of characters.
// It also returns possible error during generation.
func RandomString(n int) (string, error) {
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(alphanum)))

	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buffer[i] = alphanum[index]
	}

	return string(buffer), nil
}

func randomInt(max *big.Int) (int, error) {
	rand, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(rand.Int64()), nil
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyPhoneFormat(phone string) bool {
	return len(phone) == 11
}
