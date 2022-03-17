package data

import (
	"context"

	"xorm.io/xorm"
)

type Repository[T, ID any] struct {
	db *xorm.Engine
}

func NewRepository[T, ID any](db *xorm.Engine) *Repository[T, ID] {
	return &Repository[T, ID]{db: db}
}

func (r *Repository[T, ID]) Create(ctx context.Context, ent *T) error {
	if _, err := r.db.NewSession().
		Context(ctx).
		Insert(ent); err != nil {
		return err
	}
	return nil
}

func (r *Repository[T, ID]) GetByID(ctx context.Context, id ID) (*T, error) {
	var ent T
	has, err := r.db.NewSession().
		Context(ctx).
		Where("id = ?", id).
		Get(&ent)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &ent, nil
}

func (r *Repository[T, ID]) Count(ctx context.Context) (int64, error) {
	ent := new(T)
	count, err := r.db.NewSession().
		Context(ctx).
		Count(ent)
	if err != nil {
		return 0, err
	}
	return count, nil
}
