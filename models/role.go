package models

type role struct {
	RoleName string `xorm:"not null"`
	RoleID   int    `xorm:"not null pk"`
}
