package domain

type Article struct{}

type ArticleRepository interface {
	Insert(*Article) (*Article, error)
	Update(*Article) (*Article, error)

	FindById(int) (*Article, error)
}
