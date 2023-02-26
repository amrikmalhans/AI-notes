package audio

import (
	"context"
	"os"

	"github.com/go-kit/log"
)

type audioService struct{}

func NewService() Service {
	return &audioService{}
}

func (s *audioService) Get(ctx context.Context, id string) error {
	return nil
}

func (s *audioService) Upload(ctx context.Context, id string) error {
	return nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
