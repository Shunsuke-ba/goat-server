package domain

type Health struct {
}

func (h Health) Check() string {
	return "OK!"
}
