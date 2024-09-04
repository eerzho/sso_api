package v1

import (
	"context"
	"sso/internal/dto"
	"sso/internal/model"
)

type UserSrvc interface {
	List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.User, *dto.Pagination, error)
	Create(ctx context.Context, email, name, password string) (*model.User, error)
	Show(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, id, name string) (*model.User, error)
	Delete(ctx context.Context, id string) error
}
