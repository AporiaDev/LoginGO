package service

import (
	"errors"

	"log/internal/repository"
	"log/internal/models"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NuevoAuthRepository (repo *repository.UserRepository) *AuthService{
	return &AuthService{repo : repo}
}

func (s *AuthService) Autenticar(correo, password string)(*models.User,error){

	usuario, err := s.repo.BuscarPorCorreo(correo)

	if err != nil{
		return nil, errors.New("Credenciales Invalidas")
	}

	if usuario.Password != password{
		return nil, errors.New("Credenciales invalidas")		
	}

	return usuario, nil

}