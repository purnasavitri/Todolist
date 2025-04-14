package model

import (
	"time"
)

type Note struct {
	ID            uint           `gorm:"column:id_notes;primaryKey" json:"id"`
	Title         string         `gorm:"column:title" json:"title"`
	Content       string         `gorm:"column:content" json:"content"`
	IsFavorite    bool           `gorm:"column:is_favorite" json:"is_favorite"`
	CreatedBy     uint           `gorm:"column:created_by" json:"created_by"`
	IDCategory    uint           `gorm:"column:id_category" json:"id_category"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	User          User           `gorm:"foreignKey:id_user;references:ID" json:"user"`
	Category      Category       `gorm:"foreignKey:IDCategory;references:ID" json:"category"`
	Collaborators []Collaborator `gorm:"foreignKey:NoteID;references:ID" json:"collaborators"`
}

func (Note) TableName() string {
	return "notes"
}
