package hash

type PasswordHash interface {
	ComparePasswordAndHash(password string, encodedHash string) (match bool, err error)
	GeneratePassword(password string) (encodedHash string, err error)
}