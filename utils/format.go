package utils

import (
	m "scriptura/scriptura-api/models"
)

func FormatResponse(resource interface{}) m.SingleResponse {
	var response m.SingleResponse
	response.Data = resource
	return response
}

func FormatPaginationResponse(resources interface{}, totalItems int, offset int, limit int) (m.SliceResponse, error) {
	var response m.SliceResponse
	var err error

	response.Data = resources
	response.Pagination.CurrentPage = offset/limit + 1
	response.Pagination.TotalItems = totalItems
	response.Pagination.PageSize = limit
	response.Pagination.TotalPages = (totalItems + limit - 1) / limit

	return response, err
}
