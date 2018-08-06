package models

type authority struct {
	AutorityID   int    `xorm:"not null pk"`
	AutorityName string `xorm:"not null"`
}
