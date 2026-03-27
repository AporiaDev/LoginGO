package models

type User struct {

	//Variable tipo , mascara para Gin, mascara en base de datos

	ID int `json: "id" db: "id_usuario"`
	Correo string `json: "correo" binding: "required,email  db:"correo"`
	Password string `json: "password" binding: "required,min=8"  db: "password"`
	Nombre string `json: "nombre " db: "nombre_completa"`

}