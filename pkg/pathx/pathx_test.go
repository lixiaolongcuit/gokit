package pathx

import "testing"

func TestPathExists(t *testing.T) {
	exist, err := PathExists("pathx.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("exist:%v", exist)
}
