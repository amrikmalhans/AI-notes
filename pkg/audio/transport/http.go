package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/amrikmalhans/AI-notes/pkg/audio/endpoints"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/get", httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))

	mux.Handle("/upload", httptransport.NewServer(
		ep.UploadEndpoint,
		decodeHTTPUploadRequest,
		encodeResponse,
	))

	return mux
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.GetRequest

	// Read Query Parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, errors.New("unknown argument passed")
	}

	request.Id = id

	return request, nil
}

func decodeHTTPUploadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.UploadRequest

	// Parse the multipart/form-data
	err := r.ParseMultipartForm(32 << 20)

	if err != nil {
		return nil, errors.New("invalid argument passed")
	}

	// Extract the file from the form data
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, errors.New("invalid argument passed")
	}
	defer file.Close()

	request = endpoints.UploadRequest{
		File:   file,
		Header: *header,
	}

	return request, nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errors.New("unknown argument passed"):
		w.WriteHeader(http.StatusNotFound)
	case errors.New("invalid argument passed"):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
