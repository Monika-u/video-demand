package models

type User struct {
	// gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
