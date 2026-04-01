package models

type User struct {

	//Variable tipo , mascara para Gin, mascara en base de datos

	ID int `db: "id_usuario"`
	Correo string `db:"correo"`
	Password string `db: "password"`
	Nombre string ` db: "nombre_completa"`

}