package data

import (
	"context"
	"fmt"

	"xorm.io/xorm"

	"github.com/Jimeux/go-generic-repo/cache"
)

type Repository[T, ID any] struct {
	db  *xorm.Engine
	lru *cache.Cache[string, *T]
}

func NewRepository[T, ID any](db *xorm.Engine) *Repository[T, ID] {
	return &Repository[T, ID]{
		db:  db,
		lru: cache.NewCache[string, *T](100),
	}
}

func (r *Repository[T, ID]) Create(ctx context.Context, ent *T) error {
	_, err := r.db.NewSession().
		Context(ctx).
		Insert(ent)
	return err
}

func (r *Repository[T, ID]) GetByID(ctx context.Context, id ID) (*T, error) {
	if found, ok := r.lru.Get("id:" + fmt.Sprint(id)); ok {
		return found, nil
	}

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

func (r *Repository[T, ID]) Delete(ctx context.Context, id ID) error {
	r.lru.Delete("id:" + fmt.Sprint(id))

	_, err := r.db.NewSession().
		Context(ctx).
		Where("id = ?", id).
		Delete(new(T))
	return err
}

func (r *Repository[T, ID]) Update(ctx context.Context, id ID, t *T, cols []string) error {
	r.lru.Delete("id:" + fmt.Sprint(id))

	_, err := r.db.NewSession().
		Context(ctx).
		Cols(cols...).
		Where("id = ?", id).
		Update(t)
	return err
}

func (r *Repository[T, ID]) Count(ctx context.Context) (int64, error) {
	count, err := r.db.NewSession().
		Context(ctx).
		Count(new(T))
	if err != nil {
		return 0, err
	}
	return count, nil
}
