package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-note-app/data/request"
	"github.com/nattrio/go-note-app/data/response"
	"github.com/nattrio/go-note-app/helper"
	"github.com/nattrio/go-note-app/service"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{
		noteService: service,
	}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	CreateNoteRequest := request.CreateNoteRequest{}
	err := ctx.BodyParser(&CreateNoteRequest)
	helper.ErrorPanic(err)

	controller.noteService.Create(CreateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully create note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	err := ctx.BodyParser(&updateNoteRequest)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	updateNoteRequest.Id = id

	controller.noteService.Update(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully update note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	controller.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully delete note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	note := controller.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully find note by id",
		Data:    note,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	noteResponse := controller.noteService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully find all note",
		Data:    noteResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
