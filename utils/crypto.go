package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(s string) (string, error) {
	sBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes[:])
	return hash, nil
}

func ComparePassword(hashed string, compareTo string) error {
	compareToBytes := []byte(compareTo)
	base := []byte(hashed)
	return bcrypt.CompareHashAndPassword(base, compareToBytes)
}
