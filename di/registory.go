package di

import (
	"github.com/s-beats/go-cms/infra"
)

var r *Registory

func InitRegistory(db *infra.DB) {
	r = &Registory{
		db:         db,
		s3uploader: infra.NewS3Uploader(),
	}
}

type Registory struct {
	db         *infra.DB
	s3uploader *infra.S3Uploader
}

func GetRegistory() *Registory {
	return r
}
