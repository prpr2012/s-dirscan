package utils

import "testing"

func TestCheck(t *testing.T) {
	url := "http://129.211.73.107"
	t.Log(Check(url))
}
