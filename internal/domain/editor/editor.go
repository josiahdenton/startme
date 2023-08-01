package editor

type Editor interface {
	Open(path string) error
}
