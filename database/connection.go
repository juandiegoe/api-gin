package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


var connectionString = "bdda9a370fd168:6746185e@tcp(us-cdbr-east-04.cleardb.com)/heroku_fe9e5487f629e8d"

func ConexionBD()(conexion *sql.DB){

	conexion, err := sql.Open("mysql", connectionString)

	if err != nil{
		panic(err.Error())
	}
	return conexion
}
