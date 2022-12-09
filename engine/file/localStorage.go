package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type localStorageService struct {
	config Config
}

type Config struct {
	Path string
}

type Information struct {
	FileName string
	Size     int64
	MIMEType string
	Ext      string
}

func NewLocalStorageService(config Config) Service {
	return &localStorageService{config: config}
}

func (ths *localStorageService) Base64Upload(bufferFile []byte, fileInformation Information) (path string, err error) {
	path = fmt.Sprintf("%s/%s.%s", ths.config.Path, fileInformation.FileName, fileInformation.Ext)
	src := bytes.NewReader(bufferFile)
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return path, nil
}

func (ths *localStorageService) DeleteFile(path string) error {
	return os.Remove(path)
}

func ConstructBase64ToBuffer(file string) (bufferFile []byte, fileInformation Information, err error) {
	i := strings.Index(file, ",")
	raw := file[i+1:]

	bufferFile, err = base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, Information{}, err
	}

	f := bytes.NewReader(bufferFile)
	mimeType := http.DetectContentType(bufferFile)

	ext := strings.Split(mimeType, "/")[1]
	if strings.Contains(mimeType, "text/plain") {
		ext = "txt"
	}

	fileInformation = Information{
		FileName: fmt.Sprintf("%d", time.Now().Unix()),
		Size:     f.Size(),
		MIMEType: mimeType,
		Ext:      ext,
	}
	return
}

func ValidateMimeType(value string, expected []string) error {
	for _, expect := range expected {
		if strings.Contains(value, expect) {
			return nil
		}
	}
	return fmt.Errorf("not supported %s mime type", value)
}
