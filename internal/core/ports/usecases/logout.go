package usecases

type LogoutUseCase interface {
	Logout(token string) error
}
