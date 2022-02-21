package helpers

import (
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

var dateFormat = "2006-01-02"

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
	return t.Format(dateFormat)
}

// ConvertStringToTime ...
func ConvertStringToTime(ts string) time.Time {
	t, err := time.Parse(dateFormat, ts)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// MakeSlug ...
func MakeSlug(title string, newsID int64) string {
	// max slug length is 100 characters: 90 for title + 10 for newsID
	if len(title) > 90 {
		title = title[:90]
	}
	return fmt.Sprintf("%v-%v", slug.Make(title), newsID)
}
