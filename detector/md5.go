package detector

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func ComputeFileMD5(filePath string) (string, error) {
	if dir, _ := os.Stat(filePath); dir.IsDir() {
		return "", nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	md5 := fmt.Sprintf("%x", hash.Sum(nil))

	return md5, nil
}
