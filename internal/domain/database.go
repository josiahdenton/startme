package domain

type Db interface {
    connect(path string) error
}
