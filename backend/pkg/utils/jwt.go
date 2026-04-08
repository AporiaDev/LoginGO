package utils

import (
	"fmt"
	"log/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



var key = []byte("ClaveUltraSecreta")

func CocinaDeTokensEncoder(us *models.User) (string, error){

	claims := jwt.MapClaims{
		"sub" : us.ID,
		"exp" : time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key) 
	if err != nil{
		return "" , fmt.Errorf("error: %w" ,  err)
	}

	return tokenString, nil

}

func InspectorCocinadecoder(token string) {




}