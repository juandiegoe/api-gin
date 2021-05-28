package models

import "github.com/JuanDiegoE/api-gin/database"

func UserExists(user User) (bool,error) {
	var exist bool
	conexionEstablecida := database.ConexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM usuario WHERE email=?", user.Correo)

	if err != nil{
		return false,err
	}

	for registro.Next(){
		exist = true
	}

	return exist,err
}