package endpoints

import "mime/multipart"

type GetRequest struct {
	Id string `json:"id"`
}

type GetResponse struct {
	Err error  `json:"err,omitempty"`
	Id  string `json:"id"`
}

type UploadRequest struct {
	File   multipart.File       `json:"file"`
	Header multipart.FileHeader `json:"header"`
}

type UploadResponse struct {
	Ok  bool  `json:"ok"`
	Err error `json:"err,omitempty"`
}
