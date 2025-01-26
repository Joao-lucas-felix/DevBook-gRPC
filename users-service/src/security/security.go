package security

import "golang.org/x/crypto/bcrypt"

// Hash receives an password and returns the encrypted bytes
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
// VerifyPassword verify the password against hash and returns a error if they are not equivalent
func VerifyPassword(password, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}