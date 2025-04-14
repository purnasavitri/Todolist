package model

import (
	"time"
)

type Collaborator struct {
	ID      uint		`gorm:"column:id_collaborator;primaryKey" json:"id"`
	NoteID  uint		`gorm:"column:id_notes" json:"id_notes"`
	UserID  uint		`gorm:"column:id_user" json:"id_user"`
	AddedAt time.Time	`gorm:"column:added_at" json:"added_at"`
	Note   	Note		`gorm:"foreignKey:NoteID;references:ID" json:"notes"`
	User    User		`gorm:"foreignKey:UserID;references:ID" json:"user"`
}
func (Collaborator) TableName() string {
	return "collaborators"
}