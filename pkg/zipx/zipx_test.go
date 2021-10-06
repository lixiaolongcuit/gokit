package zipx

import (
	"testing"
)

func TestZip(t *testing.T) {
	if err := NewZipSecret("test.zip", "123456", BasePath("/home/lixiaolong/workspace/skyguard-reconsitution/gokit")).Add("config", "timex").Zip(); err != nil {
		t.Fatal(err)
	}
}

func TestUnzip(t *testing.T) {
	if err := NewZipSecret("test.zip", "123456", BasePath("abc"), AfterUnzipRem(true)).Unzip(); err != nil {
		t.Fatal(err)
	}
}
