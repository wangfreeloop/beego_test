package models

type user struct {
	ID       int    `xorm:"not null pk"`
	UserName string `xorm:"not null"`
	Email    string `xorm:"not null"`
	Account  string `xorm:"not null "`
	Password string `xorm:"not null"`
	RoleID   int    `xorm:"not null"`
	RoleName string `xorm:"not null"`
}
