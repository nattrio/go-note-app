package service

import (
	"github.com/nattrio/go-note-app/data/request"
	"github.com/nattrio/go-note-app/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
