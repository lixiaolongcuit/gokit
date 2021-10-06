package cryptox

import "testing"

func TestMd5File(t *testing.T) {
	md5Str, err := Md5File("./md5.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("md5 file result: %s", md5Str)
}

func TestMd5Bytes(t *testing.T) {
	md5Str := Md5Bytes([]byte("test123"))
	t.Logf("md5 bytes result: %s", md5Str)
}
