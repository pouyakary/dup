package detector

import (
	"dup/command"
	"fmt"
	"hash/adler32"
	"os"
	"path/filepath"
	"strings"

	exifremove "github.com/scottleedavis/go-exif-remove"
)

func ComputeFileMD5(path string, context *command.ComputationContext) (string, error) {
	if fileInfo, _ := os.Stat(path); fileInfo.IsDir() {
		return "", nil
	}

	// reading the fileData
	fileData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// removing data that interferes with the checksum
	normalizedBytes, err := normalizeBytes(path, fileData, context)
	if err != nil {
		return "", err
	}

	hash := checksum(normalizedBytes)
	return hash, nil
}

func checksum(bytes []byte) string {
	checksum := adler32.Checksum(bytes)
	hash := fmt.Sprintf("%x", checksum)
	return hash
}

func normalizeBytes(path string, bytes []byte, context *command.ComputationContext) ([]byte, error) {
	if context.Exact {
		return bytes, nil
	}

	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".jpeg", ".jpg":
		return normalizeJPEGBytes(bytes)
	default:
		return bytes, nil
	}
}

func normalizeJPEGBytes(jpeg []byte) ([]byte, error) {
	data, err := exifremove.Remove(jpeg)
	if err != nil {
		return nil, err
	}

	return data, nil
}
