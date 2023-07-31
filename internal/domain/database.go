package domain

type Db interface {
    connect(path string) error
    save(s Starter) error
    all() []Starter
}

