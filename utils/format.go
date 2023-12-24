package utils

import (
	"fmt"
	m "scriptura/scriptura-api/models"
)

func FormatResponse(resource interface{}) m.SingleResponse {
	var response m.SingleResponse
	response.Data.Attributes = resource
	return response
}

func FormatPaginationResponse(resources interface{}, totalItems int, offset int, limit int) (m.SliceResponse, error) {
	var response m.SliceResponse

	responseItems, err := formatResponseItems(resources)
	response.Data = responseItems
	response.Meta.Pagination.TotalItems = totalItems
	response.Meta.Pagination.CurrentPage = offset/limit + 1
	response.Meta.Pagination.PageSize = limit
	response.Meta.Pagination.TotalPages = (totalItems + limit - 1) / limit

	return response, err
}

func formatResponseItems(resources interface{}) ([]m.ResponseItem, error) {
	switch r := resources.(type) {
	case []m.Verse:
		responseItems := make([]m.ResponseItem, len(r))
		for i, verse := range r {
			responseItems[i] = m.ResponseItem{
				Id:         verse.Id,
				Type:       "verse",
				Attributes: verse,
			}
		}
		return responseItems, nil
	default:
		return nil, fmt.Errorf("model not supported")
	}
}
