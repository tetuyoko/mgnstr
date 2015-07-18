package mgnstr

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIncludesAny(t *testing.T) {
	assert.True(
		t,
		IncludesAny("hoge", []string{"hage", "hoge", "hug"}),
	)

	assert.False(
		t,
		IncludesAny("hoge", []string{"hage", "huge", "hug"}),
	)
}
