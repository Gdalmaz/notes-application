package models

import "time"

type Notes struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	User       User      `gorm :"foreignKey:UserID"`
	NotesTitle string    `json:"notestitle"`
	NotesText  string    `json:"notestext"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
