package rdb

import "context"

type SqlHandler interface {
	Execute(ctx context.Context, sql string, params ...interface{}) (uint, error)
	Query(obj interface{}, sql string, params ...interface{}) error
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
}
