package storage

type Storage interface {
	CloseDB() error
	GetName() string
}
