package bcrypter

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(data string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	hash = string(bytes)
	return hash, err
}

func ComparePasswordAndHash(data, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)); err != nil {
		return false
	}
	return true
}
