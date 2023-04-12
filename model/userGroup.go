package model

type UserGroup struct {
	ID              int `gorm:"primaryKey; autoIncrement"`
	CreatedAt       int
	UpdatedAt       int
	Name            string `gorm:"unique; not null"`
	UpgradeValve    int
	AdditionalIP    int
	AdditionalSpeed int
	Nodes           []*Node
	Users           []*User
}
