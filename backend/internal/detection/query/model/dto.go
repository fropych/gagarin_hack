package model

import (
	"mime/multipart"

	"gagarin/internal/shared"
)

type StreamQueryReq struct {
	Source string `json:"source" validate:"required"`
	Model  string `json:"model" validate:"required"`
}

type QueryCreate struct {
	Source string
	Type   shared.QueryType
	Video  *multipart.FileHeader
	Model  shared.ModelType
}

type QueryResponse struct {
	Id int64 `json:"id"`
}

type ResultReq struct {
	Id int64 `json:"id" validate:"required,gte=1"`
}
