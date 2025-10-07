package repo

import "context"

type extra interface {
	CalculateCount(ctx context.Context, id int64) (int64, error)
}
