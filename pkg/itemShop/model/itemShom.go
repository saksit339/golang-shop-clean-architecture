package model

type Item struct {
	Id          uint64
	Name        string
	Description string
	Picture     string
	Price       uint
	IsArchived  bool
}
