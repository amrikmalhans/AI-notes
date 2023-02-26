package transport

import (
	"context"
	"errors"

	"github.com/amrikmalhans/AI-notes/pkg/audio/endpoints"

	"github.com/amrikmalhans/AI-notes/api/protos/audio"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get    grpctransport.Handler
	upload grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) audio.AudioServer {
	return &grpcServer{
		get:    grpctransport.NewServer(ep.GetEndpoint, decodeGRPCGetRequest, decodeGRPCGetResponse),
		upload: grpctransport.NewServer(ep.UploadEndpoint, decodeGRPCUploadRequest, decodeGRPCUploadResponse),
	}
}

func (g *grpcServer) Get(ctx context.Context, req *audio.GetRequest) (*audio.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep.(*audio.GetReply), nil
}

func (g *grpcServer) Upload(ctx context.Context, req *audio.UploadRequest) (*audio.UploadReply, error) {
	_, rep, err := g.upload.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep.(*audio.UploadReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*audio.GetRequest)

	return endpoints.GetRequest{
		Id: req.TicketID,
	}, nil
}

func decodeGRPCUploadRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*audio.UploadRequest)

	return endpoints.UploadRequest{
		Id: req.TicketID,
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*audio.GetReply)

	return endpoints.GetResponse{
		Err: errors.New(reply.Err),
	}, nil
}

func decodeGRPCUploadResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*audio.UploadReply)

	return endpoints.UploadResponse{
		Err: errors.New(reply.Err),
	}, nil
}
