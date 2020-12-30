package domain

// User test
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
