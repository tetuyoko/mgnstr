package mgnstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncludesAny(t *testing.T) {
	targets := []string{"hage", "hoge", "hug"}
	assert.True(t, IncludesAny("hoge", targets))
	assert.False(t, IncludesAny("ha", targets))
}

func TestContainsAny(t *testing.T) {
	targets := []string{"__MACOSX", ".DS_Store"}
	trues := []string{"__MACOSX", "__MACOSX/hogehoge", "hoge/.DS_Store"}

	for _, sample := range trues {
		assert.True(t, ContainsAny(sample, targets))
	}

	assert.False(t, ContainsAny("hoge", targets))
}

func TestUniversalExt(t *testing.T) {
	samples := []string{"hoge.txt", "hoge.TXT"}
	for _, sample := range samples {
		assert.Equal(t, UniversalExt(sample), ".txt")
	}

	assert.NotEqual(t, UniversalExt("hoge.TXT"), ".TXT")
}

func TestGetUUID(t *testing.T) {
	sample := GetUUID()
	assert.Equal(t, len(sample), 36)
	assert.NotEqual(t, GetUUID(), sample)
}

func TestGet8UUID(t *testing.T) {
	sample := Get8UUID()
	assert.Equal(t, len(sample), 8)
	assert.NotEqual(t, Get8UUID(), sample)
}

func TestGetSha256(t *testing.T) {
	sample := GetSha256()
	sample2 := GetSha256()
	assert.Equal(t, len(sample), 32)
	assert.NotEqual(t, sample2, sample)
}
