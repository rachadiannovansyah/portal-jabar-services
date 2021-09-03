package domain

// Request ...
type Request struct {
	Keyword string
	Page    int64
	PerPage int64
	Offset  int64
	OrderBy string
	SortBy  string
}
