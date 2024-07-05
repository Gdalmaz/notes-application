package controllers

import (
	"notes-application/database"
	"notes-application/middleware"
	"notes-application/models"

	"github.com/gofiber/fiber/v2"
)

func CreateNotes(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error token step"})
	}

	notes := new(models.Notes)
	c.BodyParser(&notes)
	notestitle := c.FormValue("notestitle")
	notestext := c.FormValue("notestext")
	notes.UserID = user.ID
	if notestitle != "" {
		notes.NotesTitle = notestitle
	}
	if notestext != "" {
		notes.NotesText = notestext
	}
	err = database.DB.Db.Create(notes).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "not created notes", "data": err})
	}

	err = database.DB.Db.Preload("User").First(&notes, notes.ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error to innerjoin", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "error", "message": "created notes successfully", "data": notes})

}

func UpdateNotes(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error token step"})
	}
	notes := new(models.Notes)
	c.BodyParser(&notes)
	err = database.DB.Db.Where("user_id=?", user.ID).Find(&notes).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error user join notes step"})
	}
	err = database.DB.Db.Where("id=?", notes.ID).First(&notes).Error
	if err != nil {
		return c.Status(402).JSON(fiber.Map{"status": "error", "message": "error found notes"})
	}
	notestitle := c.FormValue("notestitle")
	notestext := c.FormValue("notestext")
	notes.NotesTitle = notestitle
	notes.NotesText = notestext
	err = database.DB.Db.Save(&notes).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "error updates step"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "your notes updated successfully"})

}

func DeleteNotes(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error token step"})
	}
	var notes models.Notes
	c.BodyParser(&notes)
	err = database.DB.Db.Where("user_id=?", user.ID).Find(&notes).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error user join notes step"})
	}
	err = database.DB.Db.Where("id=?", notes.ID).First(&notes).Error
	if err != nil {
		return c.Status(402).JSON(fiber.Map{"status": "error", "message": "error found notes"})
	}
	err = database.DB.Db.Delete(notes).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "error delete step"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "your notes delete successfully"})

}

func GetAllNotes(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error token step"})
	}
	var notes models.Notes
	err = database.DB.Db.Preload("User").Where("user_id=?", user.ID).Order("id DESC").Find(notes).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error user join notes step"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "successfull", "data": notes})

}
