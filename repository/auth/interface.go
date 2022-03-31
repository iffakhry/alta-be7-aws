package auth

type AuthRepositoryInterface interface {
	Login(email string, password string) (string, error) //login menggunakan pendekatan email dan password
}
