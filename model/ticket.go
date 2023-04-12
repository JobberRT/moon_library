package model

type Ticket struct {
	ID        int `gorm:"primaryKey; autoIncrement"`
	CreatedAt int
	UpdatedAt int
	UserID    int
	User      *User
}
