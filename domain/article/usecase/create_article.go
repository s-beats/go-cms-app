package usecase

import (
	"github.com/s-beats/go-cms/domain/article/model"
)

type UsecaseCreateArticle interface {
	Run()
}

type usecaseCreateArticle struct {
	articleRepository model.ArticleRepository
}

func NewUsecaseCreateArticle(ar model.ArticleRepository) UsecaseCreateArticle {
	return &usecaseCreateArticle{
		articleRepository: ar,
	}
}

func (u *usecaseCreateArticle) Run()
