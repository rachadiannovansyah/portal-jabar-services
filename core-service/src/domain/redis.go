package domain

// struct for redis
type Cache struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}
