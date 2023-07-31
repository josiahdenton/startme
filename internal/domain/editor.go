package domain

type Editor interface {
	Open(path string) error
	Wait() error
}
