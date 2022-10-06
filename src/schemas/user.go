package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
}
