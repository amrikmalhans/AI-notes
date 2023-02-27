package audio

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"github.com/go-kit/log"
)

type audioService struct{}

func NewService() Service {
	return &audioService{}
}

func (s *audioService) Get(ctx context.Context, id string) error {
	return nil
}

func (s *audioService) Upload(ctx context.Context, file multipart.File) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		logger.Log("errrrrrr", err)
		return err
	}

	defer client.Close()

	wc := client.Bucket("audio-files-666").Object("test").NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		logger.Log("err", err)
		return err
	}

	if err := wc.Close(); err != nil {
		logger.Log("err", err)
		return err
	}

	return nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
