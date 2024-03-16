package service

import "context"

type IStrategy interface {
	Exec(ctx context.Context) error
}
