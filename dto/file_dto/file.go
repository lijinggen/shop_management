package file_dto

import "mime/multipart"

type UploadReq struct {
	File    *multipart.FileHeader
	MaxSize int64
}
