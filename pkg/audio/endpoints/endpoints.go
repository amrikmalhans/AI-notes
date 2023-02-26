package endpoints

import (
	"context"
	"os"

	"github.com/amrikmalhans/AI-notes/pkg/audio"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Set struct {
	GetEndpoint    endpoint.Endpoint
	UploadEndpoint endpoint.Endpoint
}

func NewEndpointSet(s audio.Service) Set {
	return Set{
		GetEndpoint:    MakeGetEndpoint(s),
		UploadEndpoint: MakeUploadEndpoint(s),
	}
}

func MakeGetEndpoint(s audio.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		err := s.Get(ctx, req.Id)
		return GetResponse{Err: err}, nil
	}
}

func MakeUploadEndpoint(s audio.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UploadRequest)
		err := s.Upload(ctx, req.Id)
		return UploadResponse{Err: err}, nil
	}
}

func (s *Set) Get(ctx context.Context, id string) error {
	resp, err := s.GetEndpoint(ctx, GetRequest{Id: id})
	if err != nil {
		return err
	}
	getResp := resp.(GetResponse)
	return getResp.Err
}

func (s *Set) Upload(ctx context.Context, id string) error {
	resp, err := s.UploadEndpoint(ctx, UploadRequest{Id: id})
	if err != nil {
		return err
	}
	uploadResp := resp.(UploadResponse)
	return uploadResp.Err
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
