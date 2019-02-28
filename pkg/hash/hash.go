package hash

import "golang.org/x/crypto/bcrypt"

func In(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}

func Compare(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}
