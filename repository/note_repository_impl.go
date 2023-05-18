package repository

import (
	"errors"

	"github.com/nattrio/go-note-app/data/request"
	"github.com/nattrio/go-note-app/helper"
	"github.com/nattrio/go-note-app/model"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements NoteRepository
func (n *NoteRepositoryImpl) Delete(noteId int) {
	var note model.Note
	result := n.Db.Where("id = ?", noteId).Delete(&note)
	helper.ErrorPanic(result.Error)
}

// FindAll implements NoteRepository
func (n *NoteRepositoryImpl) FindAll() []model.Note {
	var notes []model.Note
	result := n.Db.Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}

// FindById implements NoteRepository
func (n *NoteRepositoryImpl) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := n.Db.Find(&note, noteId)
	if result != nil {
		return note, nil
	} else {
		return note, errors.New("note not found")
	}
}

// Save implements NoteRepository
func (n *NoteRepositoryImpl) Save(note model.Note) {
	result := n.Db.Create(&note)
	helper.ErrorPanic(result.Error)
}

// Update implements NoteRepository
func (n *NoteRepositoryImpl) Update(note model.Note) {
	var updateNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.Db.Model(&note).Updates(updateNote)
	helper.ErrorPanic(result.Error)
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db}
}
