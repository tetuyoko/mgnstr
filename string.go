package mgnstr

import (
	"fmt"
	"path"
	"strings"

	"crypto/rand"
	"crypto/sha256"
	"github.com/satori/go.uuid"
)

// target文字列がstrs要素として完全一致で一つでも存在する
func IncludesAny(target string, strs []string) bool {
	for _, str := range strs {
		if target == str {
			return true
		}
	}
	return false
}

// target文字列中にstrs要素のいずれかの文字列が出現する
func ContainsAny(target string, strs []string) bool {
	for _, str := range strs {
		if strings.Contains(target, str) {
			return true
		}
	}
	return false
}

// pathに含まれるextのうち、大文字、小文字を全て小文字として扱う
func UniversalExt(pth string) string {
	return strings.ToLower(path.Ext(pth))
}

// UUIDを生成する
func GetUUID() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}

// 8桁のUUIDを生成する
func Get8UUID() string {
	u1 := uuid.NewV4()
	str := fmt.Sprintf("%s", u1)
	return str[0:8]
}

// api key用の文字列を生成する
func GetSha256() (n string, err error) {
	data := make([]byte, 10)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}
	uintdata := fmt.Sprintf("%x", sha256.Sum256(data))
	return string(uintdata)[0:32], nil
}
