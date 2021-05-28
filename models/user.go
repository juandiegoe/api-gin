package models

import (
	"github.com/JuanDiegoE/api-gin/database"

	"golang.org/x/crypto/bcrypt"
)


type User struct {
	Id       int
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

func NewUser(user User) (User,error) {
	
	nombre := user.Nombre
	correo := user.Correo
	
	password,err := bcrypt.GenerateFromPassword([]byte(user.Password),14)

	stringPassword := string(password)

	if err != nil{
		return User{},err
	}

	conexionEstablecida := database.ConexionBD()
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO usuario(name,email,password) Values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec(nombre, correo, stringPassword)
	user.Password = stringPassword
	return user,nil
}

func LoginUser(user User) (User,error) {

	conexionEstablecida := database.ConexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM usuario WHERE email=?", user.Correo)

	if err != nil{
		return user,err
	}

	var(
		idDB int
		nameDB string
		emailDB string
		passwordDB string
	)

	for registro.Next(){
		
		err = registro.Scan(&idDB,&nameDB,&emailDB,&passwordDB)

		if err != nil{
			return user,err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordDB),[]byte(user.Password)); err != nil{
		return user,err
	} else {
		pass,_ := bcrypt.GenerateFromPassword([]byte(passwordDB),14)
		user.Id =  idDB
		user.Nombre = nameDB
		user.Password = string(pass)
		return user,nil
	}
}

