package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log/internal/service"
)

type AuthHandler struct{
	svc *service.AuthService
}

func NuevoAuthHandler(svc *service.AuthService)*AuthHandler{
	return &AuthHandler{svc:svc}
}

//Estructura para recibir el JSON
type LoginInput struct{
	Correo string `json : "correo" bidding:"required,email"`
	Password string `json: "password" bidding:"required"`
}

func (h  *AuthHandler) Login(c *gin.Context){
	
	//variable para guardar los valores del input
	var input LoginInput
// 1. La Recepcionista lee el JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// Asume que tienes tus funciones Fail() y OK() disponibles
		Fail(c, http.StatusBadRequest, "DATOS_INVALIDOS", "Formato incorrecto")
		return
	}

	

}