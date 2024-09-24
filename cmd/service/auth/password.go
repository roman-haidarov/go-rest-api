package auth

import (
	"github.com/roman-haidarov/go-rest-api/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	data := []byte(config.Envs.ApiKey + password)
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
