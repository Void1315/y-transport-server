package model

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func CheckAuth(phone string, password string) bool {
	var auth Auth
	Db.Select("id").Where(Auth{Phone: phone, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func GetJWTUser() {

}
