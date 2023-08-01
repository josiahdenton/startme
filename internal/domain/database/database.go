package database

import (
	"github.com/josiahdenton/startme/internal/domain/starter"
)

type Db interface {
	Connect(path string) error
	Setup()
	Close() 
	Save(title, content string) error
	All() []starter.Starter
}
