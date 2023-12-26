package models

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

type ResponseItemLinks struct {
	Self        string `json:"self,omitempty"`
	Related     string `json:"related,omitempty"`
	DescribedBy string `json:"describedby,omitempty"`
}
