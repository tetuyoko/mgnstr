package mgnstr

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := Sum(10, 20)
	expected := 30
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
