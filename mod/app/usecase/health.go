package usecase

import (
	"context"

	"github.com/Shunsuke-ba/goat-server/mod/app/domain"
)

type Health func(ctx context.Context) string

func ProvideHealthCase(
	ctx context.Context,
	repo domain.Health,
) Health {

	return func(ctx context.Context) string {
		return repo.Check()
	}
}
