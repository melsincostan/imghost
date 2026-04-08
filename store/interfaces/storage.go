package interfaces

import "io"

type Storage interface {
	Store(name string, contents io.Reader) (err error)
	Retrieve(name string) (contents io.Reader, err error)
	Exists(name string) (err error)
}
