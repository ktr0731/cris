package ports

import "io"

type StoragePort interface {
	Upload(name string, content io.Reader) (string, error)
	Download(name string) (io.ReadCloser, error)
}
