package utils

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func VerifyPassword(hashedPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))

	err2 := scrypt.Key([]byte(hashedPassword))
	return err == nil
}

func PasswordScrypt(plainPassword []byte) []byte {
	salt := make([]byte, 32)
	rand.Read(salt)
	passScrypt, _ := scrypt.Key(plainPassword, salt, 1<<15, 8, 1, 32)
	return passScrypt
}
