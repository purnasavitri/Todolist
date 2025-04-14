package controller

import (
	"net/http"
	"strconv"
	"todolist/config"
	"todolist/model"

	"github.com/gin-gonic/gin"
)

func GetAllNotes(c *gin.Context) {
	var notes []model.Note
	if err := config.DB.Preload("User").Preload("Category").Preload("Collaborators").Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func GetNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id_notes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var note model.Note
	if err := config.DB.Preload("User").Preload("Category").Preload("Collaborators").First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

func CreateNote(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func UpdateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id_notes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var note model.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id_notes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var note model.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if err := config.DB.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

func GetNotesByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id_category"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var notes []model.Note
	if err := config.DB.Preload("User").Preload("Category").Preload("Collaborators").Where("id_category = ?", categoryID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes by category"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func GetFavoriteNotes(c *gin.Context) {
	var notes []model.Note
	if err := config.DB.Preload("User").Preload("Category").Preload("Collaborators").Where("is_favorite = ?", true).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorite notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}
