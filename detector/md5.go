package detector

import (
	"fmt"
	"hash/adler32"
	"os"
)

func ComputeFileMD5(filePath string) (string, error) {
	if dir, _ := os.Stat(filePath); dir.IsDir() {
		return "", nil
	}

	file, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	checksum := adler32.Checksum(file)
	hash := fmt.Sprintf("%x", checksum)

	return hash, nil
}
