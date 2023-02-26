package audio

import "context"

type Service interface {
	Get(ctx context.Context, id string) error
	Upload(ctx context.Context, id string) error
}
