package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		log.Println("error hashing the password")
		return "", err
	}
	return string(hashedpassword), nil
}

func ComparePassword(storedpassword, providedpassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(providedpassword))
	return err == nil
}