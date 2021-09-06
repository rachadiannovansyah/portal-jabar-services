package domain

// ErrorResults ...
type ErrorResults struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResultData ...
type ResultData struct {
	Data interface{} `json:"data"`
}

// ResultsData ...
type ResultsData struct {
	Data  interface{} `json:"data"`
	Links *LinksData  `json:"links"`
	Meta  *MetaData   `json:"meta"`
}

// MetaData ...
type MetaData struct {
	TotalCount  int64   `json:"total_count"`
	TotalPage   float64 `json:"total_page"`
	CurrentPage int64   `json:"current_page"`
	PerPage     int64   `json:"per_page"`
}

// LinksData ...
type LinksData struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
	Next  *string `json:"next"`
	Prev  *string `json:"prev"`
}
