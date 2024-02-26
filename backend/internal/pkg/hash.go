package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(psw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(psw), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
