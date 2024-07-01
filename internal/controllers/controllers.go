package controllers

type UseCases interface {
	EvaluatePassword(password string) int
}

type Controllers struct {
	useCases UseCases
}

func New(useCases UseCases) *Controllers {
	return &Controllers{
		useCases: useCases,
	}
}
