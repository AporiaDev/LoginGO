package repository

import (
	"database/sql"
	"errors"
	"log"
	

	"log/internal/models"
)

type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{db:db}
}

func (r *UserRepository) BuscarPorCorreo(correoBuscado string) (*models.User, error){

	query := `SELECT id, username, password, correo
				from users WHERE correo = $1`
	
	var usuario models.User
	
	//pedir comida a domicilio
	fila := r.db.QueryRow(query, correoBuscado)

	//con scan ya me puedo comer lo que pedi a domicilio
	//scan me ayuda a hacer el parseo manual
	err := fila.Scan(

		&usuario.ID,
		&usuario.Nombre,
		&usuario.Password,
		&usuario.Correo,
		
		
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Este es un error esperado: El usuario no existe en la base de datos
			return nil, errors.New("usuario no encontrado")
		}
		// Este es un error crítico (conexión caída, tabla no existe, etc.)
		log.Printf("Error grave en la base de datos al buscar usuario: %v", err)
		return nil, err
	}

	return &usuario, nil
}




