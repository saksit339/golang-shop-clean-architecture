package model

type Item struct {
	Id          uint64
	Name        string
	Description string
	Picture     string
	Price       uint
	IsArchived  bool
}

type ItemFilter struct {
	Name        string `query:"name" validate:"omitempty,max=64"`
	Description string `query:"description" validate:"omitempty,max=128"`
}
