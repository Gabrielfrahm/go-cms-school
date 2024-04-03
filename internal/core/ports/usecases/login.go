package usecases

type LoginResponse struct {
	Token         string
	Refresh_token string
}

type LoginUseCase interface {
	Login(email, password string) (*LoginResponse, error)
}
