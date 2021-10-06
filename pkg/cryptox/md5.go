package cryptox

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//计算文件md5
func Md5File(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5 := md5.New()
	if _, err := io.Copy(md5, f); err != nil {
		return "", err
	}
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str, nil
}

func Md5Bytes(bytes []byte) string {
	md5Bytes := md5.Sum(bytes)
	md5Str := hex.EncodeToString(md5Bytes[:])
	return md5Str
}

func Md5Str(str string) string {
	return Md5Bytes([]byte(str))
}
