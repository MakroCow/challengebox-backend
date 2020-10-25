package models

type Tag struct {
	Id   int
	Name string
}

func NewTag(id int, name string) *Tag {
	return &Tag{Id: id, Name: name}
}
