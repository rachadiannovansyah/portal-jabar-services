package domain

import (
	"encoding/json"
)

// JSONStringSlices MarshalJSON for JsonStringSlices
type JSONStringSlices []string

func unmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Scan implements the JSONStringSlices Scanner interface.
func (s *JSONStringSlices) Scan(value interface{}) error {
	return unmarshalJSON(value.([]byte), s)
}

// Scan implements the SocialMedia Scanner struct.
func (s *SocialMedia) Scan(value interface{}) error {
	return unmarshalJSON(value.([]byte), s)
}

// Scan implements the PosterInfo Scanner interface.
func (s *PosterInfo) Scan(value interface{}) error {
	return unmarshalJSON(value.([]byte), s)
}
