package mgnstr

import (
	"fmt"
	"path"
	"strings"

	"crypto/rand"
	"crypto/sha256"
	"github.com/satori/go.uuid"
)

// target string is also exist in one for an exact match as strs element
func IncludesAny(target string, strs []string) bool {
	for _, str := range strs {
		if target == str {
			return true
		}
	}
	return false
}

// one of the strings of strs elements appear in the target string
func ContainsAny(target string, strs []string) bool {
	for _, str := range strs {
		if strings.Contains(target, str) {
			return true
		}
	}
	return false
}

// Of the ext that is included in the path,
// it treats all letters as lowercase
func UniversalExt(pth string) string {
	return strings.ToLower(path.Ext(pth))
}

// generate UUID v4 as string.
func GetUUID() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}

// generate UUID v4, 8 letter as string.
func Get8UUID() string {
	u1 := uuid.NewV4()
	str := fmt.Sprintf("%s", u1)
	return str[0:8]
}

// generate sha256 key
func GetSha256() (n string, err error) {
	data := make([]byte, 10)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}
	uintdata := fmt.Sprintf("%x", sha256.Sum256(data))
	return string(uintdata)[0:32], nil
}
