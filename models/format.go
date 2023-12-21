package models

type Response struct {
	Errors map[string]any    `json:"errors,omitempty"`
	Meta   map[string]any    `json:"meta,omitempty"`
	Links  ResponseItemLinks `json:"links,omitempty"`
	// Page       int               `json:"page,omitempty"`
	// PageSize   int               `json:"pageSize,omitempty"`
	// LastPage   int               `json:"lastPage,omitempty"`
	// TotalItems int               `json:"totalItems,omitempty"`
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
