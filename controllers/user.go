package controllers

import (
	"notes-application/database"
	"notes-application/helpers"
	"notes-application/middleware"
	"notes-application/models"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error to parsing step"})
	}

	err = helpers.MailControl(user.Mail)
	if err == nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "you have this account", "data": err})
	}

	err = database.DB.Db.Create(user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Fail Creating Account", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Account was created successfully", "data": user})
}

func LogIn(c *fiber.Ctx) error {
	user := new(models.User)
	Loguser := new(models.LogInInfo)
	c.BodyParser(&Loguser)
	err := database.DB.Db.Where("mail=? and password=?", Loguser.Mail, Loguser.Password).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "wrong password/mail", "data": err})
	}
	token, err := middleware.CreateToken(user.FirstName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("token oluşturulamadı")
	}
	session := new(models.Session)
	session.UserID = user.ID
	session.Token = token
	err = database.DB.Db.Create(&session).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to login", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login Success", "data": user, "token": token})
}

func UpdatePassword(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "not found to token"})
	}
	UpdateAccount := new(models.UpdateAccount)
	err = c.BodyParser(&UpdateAccount)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "fail to bodyparser step"})
	}

	UpdateAccount.OldPassword = user.Password

	if UpdateAccount.OldPassword == UpdateAccount.NewPassword1 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "your write sample password", "data": err})
	}

	if UpdateAccount.NewPassword1 != UpdateAccount.NewPassword2 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "not equal old / new", "data": err})
	}

	user.Password = UpdateAccount.NewPassword1
	user.FirstName = UpdateAccount.NewFirstName
	user.LastName = UpdateAccount.NewLastName

	err = database.DB.Db.Updates(user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error to update user step", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "successfull update password user", "data": user})
}

func LogOut(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err})
	}
	UserID := user.ID
	session := new(models.Session)

	err = database.DB.Db.Where("userid=?", UserID).First(&session).Error
	if err == nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "not found token"})
	}
	//burada OAREM YAPISI İŞ GÖRMEDİ
	err = database.DB.Db.Raw("DELETE FROM sessions WHERE user_id= ?", UserID).Scan(&session).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "not found token1"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "logout successfully"})
}

func DeleteAccount(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "not found token", "data": err})
	}
	err = database.DB.Db.Delete(user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error delete step", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "delete success"})
}
