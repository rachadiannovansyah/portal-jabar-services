package helpers

import (
	"fmt"
	"regexp"
	"strings"
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
func ConvertTimeToString(t time.Time) *string {
	timeString := t.Format(dateFormat)
	return &timeString
}

// ConvertStringToTime ...
func ConvertStringToTime(ts *string) time.Time {
	t, err := time.Parse(dateFormat, *ts)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// MakeSlug ...
func MakeSlug(title string, newsID int64) string {
	// max slug length is 100 characters: 89 for title + 1 for delimiter + 10 for newsID
	forTitleLength := 89
	if len(title) > forTitleLength {
		title = title[:forTitleLength]
	}
	return fmt.Sprintf("%v-%v", slug.Make(title), newsID)
}

// ConvertSliceToString ...
func ConverSliceToString(slice []string, delimiter string) string {
	return strings.Join(slice, delimiter)
}

func Substr(s string, n int) string {
	if len(s) > n {
		s = s[:n]
	}

	return s
}

// Generate Slug
func SlugGenerator(str string, identifier int64) string {
	strLower := strings.Fields(strings.ToLower(str))
	strSlug := strings.Join(strLower, "-")
	return fmt.Sprintf("%s-%d", strSlug, identifier)
}

func SlugFormatter(str string) string {
	// Convert the text to lowercase
	slug := strings.ToLower(str)

	// Remove unwanted characters using regular expressions
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading and trailing
	slug = strings.Trim(slug, "-")

	return slug
}
