package domain

type HealthRepository struct {
}

func (h HealthRepository) Check() string {
	return "OK!"
}
