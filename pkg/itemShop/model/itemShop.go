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
	Pagination
}

type Pagination struct {
	Page int64 `query:"page" validate:"required,min=1"`
	Size int64 `query:"size" validate:"required,min=1,max=100"`
}

type ItemResult struct {
	Items      []Item
	Pagination PaginationResult
}

type PaginationResult struct {
	Page      int64
	TotalPage int64
}
