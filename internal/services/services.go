package services

type services struct {
	svc Services
}

type Services interface {
}

func NewServices(s Services) Services {
	return &services{svc: s}
}
