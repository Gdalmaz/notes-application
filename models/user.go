package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Mail      string `json:"mail"`
}

type LogInInfo struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type UpdateAccount struct {
	OldPassword  string `json:"oldpassword"`
	NewPassword1 string `json:"newpassword1"`
	NewPassword2 string `json:"newpassword2"`
	NewFirstName string `json:"newfirstname"`
	NewLastName  string `json:"newlastname"`
	NewMail      string `json:"newmail"`
}
