package dto

type ParamList struct {
	Page       int    `json:"page" valid:"Required"`
	PerPage    int    `json:"per_page" valid:"Required"`
	Search     string `json:"search,omitempty"`
	InitSearch string `json:"init_search,omitempty"`
	SortField  string `json:"sort_field,omitempty"`
}
