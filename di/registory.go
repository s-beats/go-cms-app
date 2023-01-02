package di

import (
	"github.com/s-beats/go-cms/infra/db"
	"github.com/s-beats/go-cms/infra/storage"
)

var r *Registory

func InitRegistory(db *db.DB) {
	r = &Registory{
		db:         db,
		s3uploader: storage.NewS3Uploader(),
	}
}

type Registory struct {
	db         *db.DB
	s3uploader *storage.S3Uploader
}

func GetRegistory() *Registory {
	return r
}
