package models

import (
	"apirestGorm/db"
)

type User struct { //remplazamos el json por xml o yaml si queremos cambiar el formato devuelto
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//Users lista de User
type Users []User

func MigrarUser() {
	db.Database.AutoMigrate(User{})
}
