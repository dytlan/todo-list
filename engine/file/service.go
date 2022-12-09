package file

type Service interface {
	Base64Upload(bufferFile []byte, fileInformation Information) (path string, err error)
	DeleteFile(path string) error
}
