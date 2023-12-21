package models

type Response struct {
	Errors map[string]any    `json:"errors,omitempty"`
	Meta   ResponseMeta      `json:"meta,omitempty"`
	Links  ResponseItemLinks `json:"links,omitempty"`
}

type SliceResponse struct {
	Data []ResponseItem `json:"data,omitempty"`
	Response
}

type SingleResponse struct {
	Data ResponseItem `json:"data,omitempty"`
	Response
}

type ResponseItem struct {
	Type          string                  `json:"type,omitempty"`
	Id            string                  `json:"id,omitempty"`
	Attributes    interface{}             `json:"attributes,omitempty"`
	Relationships map[string]ResponseItem `json:"relationships,omitempty"`
	Links         ResponseItemLinks       `json:"links,omitempty"`
}

type ResponseItemLinks struct {
	Self        string `json:"self,omitempty"`
	Related     string `json:"related,omitempty"`
	DescribedBy string `json:"describedby,omitempty"`
}

type ResponseMeta struct {
	Pagination struct {
		TotalCount  int64 `json:"totalItems,omitempty"`
		CurrentPage int64 `json:"currentPage,omitempty"`
	} `json:"pagination,omitempty"`
}
