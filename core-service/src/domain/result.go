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
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

// MetaData ...
type MetaData struct {
	TotalCount  int64   `json:"totalCount"`
	TotalPage   float64 `json:"totalPage"`
	CurrentPage int64   `json:"currentPage"`
	PerPage     int64   `json:"perPage"`
}
