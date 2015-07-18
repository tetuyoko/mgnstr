package mgnstr

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIncludesAny(t *testing.T) {
	targets := []string{"hage", "hoge", "hug"}
	assert.True(t,  IncludesAny("hoge", targets))
	assert.False(t, IncludesAny("ha",   targets))
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
		assert.Equal(t,  UniversalExt(sample), ".txt")
	}

	assert.NotEqual(t,  UniversalExt("hoge.TXT"), ".TXT")
}
