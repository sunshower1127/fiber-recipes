package database

import "gorm.io/gorm"

type Quote struct {
	Author     string `json:"author"`
	Text       string `json:"quote"`
	gorm.Model        // 상속이라고 보면 됨 대충. 프로퍼티는 이렇게 구현한다~
}

type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	URL         string `json:"url"`
	Rating      string `json:"rating"`
	gorm.Model
}
