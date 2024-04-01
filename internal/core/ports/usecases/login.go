package usecases

type LoginResponse struct {
	Token string
}

type LoginUseCase interface {
	Login(email, password string) (*LoginResponse, error)
}
