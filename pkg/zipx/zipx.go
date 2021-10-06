package zipx

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/yeka/zip"
)

var ErrZipFilesEmpty = errors.New("zip files empty error")

type Option func(*zipSecret)

//指定了basePath后调用Add时的相对路径则相对于basePath
func BasePath(basePath string) Option {
	return func(z *zipSecret) {
		z.basePath = basePath
	}
}

//加密方法
func EncryptionMethod(method zip.EncryptionMethod) Option {
	return func(z *zipSecret) {
		z.encryptionMethod = method
	}
}

//解压后是否删除zip文件
func AfterUnzipRem(rem bool) Option {
	return func(z *zipSecret) {
		z.afterUnzipRem = rem
	}
}

type zipSecret struct {
	zipFile          string
	password         string
	basePath         string
	encryptionMethod zip.EncryptionMethod
	afterUnzipRem    bool
	files            []string
}

func NewZipSecret(zipFile, passowrd string, opts ...Option) *zipSecret {
	z := &zipSecret{
		zipFile:          zipFile,
		password:         passowrd,
		encryptionMethod: zip.StandardEncryption,
	}
	for _, op := range opts {
		op(z)
	}
	return z
}

func (z *zipSecret) Add(files ...string) *zipSecret {
	z.files = append(z.files, files...)
	return z
}

//压缩文件
func (z *zipSecret) Zip() (err error) {
	if len(z.files) == 0 {
		return ErrZipFilesEmpty
	}
	zipf, err := os.Create(z.zipFile)
	if err != nil {
		return fmt.Errorf("create zip file error %w", err)
	}
	// 压缩过程中出错删除中间文件
	defer func() {
		if err != nil {
			err = os.Remove(z.zipFile)
		}
	}()
	defer zipf.Close()
	zipWriter := zip.NewWriter(zipf)
	defer zipWriter.Close()

	for len(z.files) > 0 {
		file := z.files[0]
		z.files = z.files[1:]
		filePath := file
		if !strings.HasPrefix(file, "/") && z.basePath != "" {
			filePath = path.Join(z.basePath, filePath)
		}
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}
		//处理目录
		if stat.IsDir() {
			subFiles, err := ioutil.ReadDir(filePath)
			if err != nil {
				return err
			}
			for _, f := range subFiles {
				z.files = append(z.files, path.Join(file, f.Name()))
			}
			continue
		}
		//处理文件
		if err := z.addFile(zipWriter, file, filePath); err != nil {
			return err
		}
	}
	err = zipWriter.Flush()

	return err
}

func (z *zipSecret) addFile(zipWriter *zip.Writer, inZipFileName, filePath string) error {
	w, err := zipWriter.Encrypt(inZipFileName, z.password, z.encryptionMethod)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := io.Copy(w, file); err != nil {
		return err
	}
	return nil
}

//解压文件
func (z *zipSecret) Unzip() error {
	zipReader, err := zip.OpenReader(z.zipFile)
	if err != nil {
		return err
	}
	defer func() {
		if z.afterUnzipRem {
			os.Remove(z.zipFile)
		}
	}()
	defer zipReader.Close()
	if z.basePath != "" {
		if err = os.MkdirAll(z.basePath, os.ModePerm); err != nil {
			return err
		}
	}
	for _, zf := range zipReader.File {
		if err = z.unzipFile(zf); err != nil {
			return err
		}
	}
	return nil
}

func (z *zipSecret) unzipFile(zf *zip.File) error {
	zf.SetPassword(z.password)
	zfReader, err := zf.Open()
	if err != nil {
		return err
	}
	defer zfReader.Close()
	filePath := path.Join(z.basePath, zf.Name)
	if err = os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	fout, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fout.Close()
	_, err = io.Copy(fout, zfReader)

	return err
}
