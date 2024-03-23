package models

type User struct {
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255)"`
}
