package stringutils

import (
	"strings"
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
)

func EqualsIgnoreCase(s1 string,s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func Md5(src string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(src))
	cipherStr := md5Ctx.Sum([]byte(nil))
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}

func UUID() string  {
	id := uuid.NewV4()
	return id.String()
}