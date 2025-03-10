package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
}

// struct에 메타데이터를 이렇게 추가함.
