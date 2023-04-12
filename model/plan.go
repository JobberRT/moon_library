package model

type Plan struct {
	ID           int `gorm:"primaryKey; autoIncrement"`
	CreatedAt    int
	UpdatedAt    int
	Name         string `gorm:"unique; not null"`
	Type         string `gorm:"unique; not null"`
	Bandwidth    int
	Device       int
	Speed        int
	EnableCoupon bool `gorm:"not null"`
	Description  string
	Users        []*User
}
