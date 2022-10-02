package usecase

import "github.com/s-beats/go-cms/domain_model/domain"

type UsecaseCreateArticle interface {
	Run()
}

type usecaseCreateArticle struct {
	articleRepository domain.ArticleRepository
}

func NewUsecaseCreateArticle(ar domain.ArticleRepository) UsecaseCreateArticle {
	return &usecaseCreateArticle{
		articleRepository: ar,
	}
}

func (u *usecaseCreateArticle) Run()
