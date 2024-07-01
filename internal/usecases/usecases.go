package usecases

type PasswordService interface {
	CalculateVariety(password string) int
	CalculateEntropy(password string) int
	CalculatePoolSize(password string) int
	CheckCommonPatterns(password string) int
}

type UseCases struct {
	password PasswordService
}

func New(passwordService PasswordService) *UseCases {
	return &UseCases{
		password: passwordService,
	}
}
