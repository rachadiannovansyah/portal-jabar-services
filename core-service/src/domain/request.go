package domain

// Request ...
type Request struct {
	Keyword string
	PerPage int64
	Offset  int64
	OrderBy string
	SortBy  string
}
