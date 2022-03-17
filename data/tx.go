package data

import (
	"context"

	"xorm.io/xorm"
)

func Tx[T any](ctx context.Context, db *xorm.Engine, fn func(ctx context.Context) (T, error)) (T, error) {
	session := db.NewSession().Context(ctx)
	defer session.Close()
	var t T

	if err := session.Begin(); err != nil {
		return t, err
	}

	res, err := fn(ctx)
	if err != nil {
		_ = session.Rollback()
		return t, err
	}

	if err := session.Commit(); err != nil {
		return t, err
	}
	return res, err
}
