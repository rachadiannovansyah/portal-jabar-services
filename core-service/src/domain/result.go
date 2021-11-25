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
	TotalCount   int64             `json:"total_count"`
	TotalPage    float64           `json:"total_page"`
	CurrentPage  int64             `json:"current_page"`
	PerPage      int64             `json:"per_page"`
	Aggregations *MetaAggregations `json:"aggregations,omitempty"`
}

// CustomMetaData ..
type CustomMetaData struct {
	TotalCount  int64  `json:"total_count"`
	LastUpdated string `json:"last_updated"`
}

type MetaAggregations struct {
	Domain AggDomain `json:"domain"`
}

// AggDomain
type AggDomain struct {
	News          int64 `json:"news"`
	Information   int64 `json:"information"`
	PublicService int64 `json:"public_service"`
	Announcement  int64 `json:"announcement"`
	About         int64 `json:"about"`
}
