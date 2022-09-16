package factory

import (
	"github.com/s-beats/go-cms/infra"
)

type Registory struct {
	db *infra.DB
	s3 *infra.S3Uploader
}

func NewRegistory(db *infra.DB, s3 *infra.S3Uploader) *Registory {
	return &Registory{
		db: db,
		s3: s3,
	}
}
