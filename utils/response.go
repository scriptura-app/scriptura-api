package utils

import (
	"encoding/json"
	"os"
	"strconv"
)

type Response struct {
}

type SliceResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination struct {
		Total  int `json:"total,omitempty"`
		Offset int `json:"offset,omitempty"`
		Cursor int `json:"cursor,omitempty"`
		Limit  int `json:"limit"`
	} `json:"pagination,omitempty"`
	Response
}

type SingleResponse struct {
	Data interface{} `json:"data,omitempty"`
	Response
}

func FormatResponse(resource interface{}) []byte {
	var response SingleResponse
	response.Data = resource
	out, _ := json.Marshal(response)
	return out
}

func FormatPaginationResponse(resources interface{}, total int, offset int, limit int, cursor int) []byte {
	var response SliceResponse
	response.Data = resources
	response.Pagination.Total = total
	response.Pagination.Offset = offset
	response.Pagination.Limit = limit
	response.Pagination.Cursor = cursor
	out, _ := json.Marshal(response)
	return out
}

func GetURI(entity string, id int) string {
	return os.Getenv("BASE_URL") + "/api/v1/" + entity + "/" + strconv.Itoa(id)
}
