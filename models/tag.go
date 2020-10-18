package models

type Tag struct {
	id   int
	name string
}

func NewTag(id int, name string) *Tag {
	return &Tag{id: id, name: name}
}
