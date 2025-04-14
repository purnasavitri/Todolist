package controller

import (
	"net/http"
	"strconv"
	"todolist/config"
	"todolist/model"

	"github.com/gin-gonic/gin"
)

func GetCollaboratorsByNoteID(c *gin.Context) {
	noteID, err := strconv.Atoi(c.Param("id_notes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var collaborators []model.Collaborator
	if err := config.DB.Where("id_notes = ?", noteID).Preload("Note").Preload("User").Find(&collaborators).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch collaborators"})
		return
	}
	c.JSON(http.StatusOK, collaborators)
}

func AddCollaborator(c *gin.Context) {
	noteID, errNote := strconv.Atoi(c.Param("id_notes"))
	userID, errUser := strconv.Atoi(c.Param("id_user"))
	if (errNote != nil) || (errUser != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID or user ID"})
		return
	}

	var existingCollaborator model.Collaborator
    err := config.DB.
        Where("id_notes = ? AND id_user = ?", noteID, userID).
        First(&existingCollaborator).Error

    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User is already a collaborator in this note"})
        return
    }

	collaborator := model.Collaborator{NoteID: uint(noteID), UserID: uint(userID)}

	if err := config.DB.Create(&collaborator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add collaborator"})
		return
	}
	c.JSON(http.StatusOK, collaborator)
}

func RemoveCollaborator(c *gin.Context) {
	noteID, errNote := strconv.Atoi(c.Param("id_notes"))
	userID, errUser := strconv.Atoi(c.Param("id_user"))
	if (errNote != nil) || (errUser != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID or user ID"})
		return
	}

	if err := config.DB.Where("id_notes = ? AND id_user = ?", noteID, userID).Delete(&model.Collaborator{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove collaborator"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Collaborator removed successfully"})
}
