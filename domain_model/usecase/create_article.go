package usecase

type UsecaseCreateArticle interface {
	Run()
}

type usecaseCreateArticle struct{}

func NewUsecaseCreateArticle() UsecaseCreateArticle {
	return &usecaseCreateArticle{}
}

func (u *usecaseCreateArticle) Run()
