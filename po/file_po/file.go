package file_po

import "mime/multipart"

type UploadResp struct {
	Url string `json:"url"`
}

type UploadReq struct {
	File    *multipart.FileHeader
	MaxSize int64
}
