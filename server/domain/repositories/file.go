package repositories

import "github.com/ktr0731/cris/server/domain/entities"

type FileRepository interface {
	Store(*entities.File) (entities.FileID, error)
	Find(entities.FileID) (*entities.File, error)
	Remove(entities.FileID) (*entities.File, error)
}
