package model

type Document struct {
	ID      int    `gorm:"primaryKey; autoIncrement"`
	Name    string `gorm:"unique; not null"`
	Content string
}
