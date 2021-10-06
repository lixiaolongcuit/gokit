package cryptox

import "testing"

func TestSha256Str(t *testing.T) {
	res := Sha256Str("test")
	t.Log(res)
}
