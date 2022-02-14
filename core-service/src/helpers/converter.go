package helpers

import "time"

// SetPointerString ...
func SetPointerString(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

// SetPointerInt64 ...
func SetPointerInt64(val int64) *int64 {
	if val == 0 {
		return nil
	}
	return &val
}

// ConvertTimeToString ...
func ConvertTimeToString(t time.Time) string {
	return t.Format("2006-01-02")
}
