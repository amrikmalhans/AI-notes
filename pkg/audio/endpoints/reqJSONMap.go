package endpoints

type GetRequest struct {
	Id string `json:"id"`
}

type GetResponse struct {
	Err error `json:"err,omitempty"`
}

type UploadRequest struct {
	Id string `json:"id"`
}

type UploadResponse struct {
	Err error `json:"err,omitempty"`
}
