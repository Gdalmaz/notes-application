package helpers

import (
	"notes-application/database"
	"notes-application/models"
)

func MailControl(mail string) error {
	user := new(models.User)
	err := database.DB.Db.Where("mail=?",mail).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}