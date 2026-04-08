package interfaces

type Migrator interface {
	Up() (applied int, err error)
}
