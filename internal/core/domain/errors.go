package domain

import "errors"

var (
	ErrInvalidPath  = errors.New("invalid file path")
	ErrInvalidCID   = errors.New("invalid CID")
	ErrFileNotFound = errors.New("file not found")
	ErrUploadFailed = errors.New("file upload failed")
)
