package database

import (
	"github.com/josiahdenton/startme/internal/domain/starter"
)

type Db interface {
	Connect(path string) error
	Setup()
	Close() 
	Save(s starter.Starter) error
	All() []starter.Starter
}
