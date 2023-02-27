package audio

import (
	"context"
	"mime/multipart"
)

type Service interface {
	Get(ctx context.Context, id string) error
	Upload(ctx context.Context, file multipart.File) error
}
