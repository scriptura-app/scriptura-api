package utils

import (
	"os"
	"strconv"
)

type Response struct {
}

type SliceResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination struct {
		TotalItems  int `json:"totalItems,omitempty"`
		CurrentPage int `json:"currentPage,omitempty"`
		PageSize    int `json:"pageSize,omitempty"`
		TotalPages  int `json:"totalPages,omitempty"`
	} `json:"pagination,omitempty"`
	Response
}

type SingleResponse struct {
	Data interface{} `json:"data,omitempty"`
	Response
}

func FormatResponse(resource interface{}) SingleResponse {
	var response SingleResponse
	response.Data = resource
	return response
}

func FormatPaginationResponse(resources interface{}, totalItems int, offset int, limit int) (SliceResponse, error) {
	var response SliceResponse
	var err error

	response.Data = resources
	response.Pagination.CurrentPage = offset/limit + 1
	response.Pagination.TotalItems = totalItems
	response.Pagination.PageSize = limit
	response.Pagination.TotalPages = (totalItems + limit - 1) / limit

	return response, err
}

func GetURI(entity string, id int) string {
	return os.Getenv("BASE_URL") + "/api/v1/" + entity + "/" + strconv.Itoa(id)
}
