package dto

import "mime/multipart"

type VideoUploadRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
