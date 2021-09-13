package domain

// Request ...
type Request struct {
	Keyword   string
	Page      int64
	PerPage   int64
	Offset    int64
	SortBy    string
	SortOrder string
	StartDate string
	EndDate   string
}
