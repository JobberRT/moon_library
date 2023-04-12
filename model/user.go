package model

type User struct {
	ID          int `gorm:"primaryKey; autoIncrement"`
	CreatedAt   int
	UpdatedAt   int
	UserGroupID int
	PlanID      int
	Username    string `gorm:"unique; not null"`
	Email       string `gorm:"unique; not null"`
	Balance     int
	Password    string `gorm:"not null"`
	ExpireAt    int64
	IsAdmin     bool `gorm:"not null"`
}
