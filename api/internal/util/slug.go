package util

import (
	"fmt"

	"github.com/gosimple/slug"
)

// TODO: move all slug generation to the database
// Slugify removes spaces and converts to lower case
func Slugify(name string) string {
	slug := slug.Make(name)
	return slug
}

// NextUniqueSlug returns the next unique slug available based on
func NextUniqueSlug(name string, slugsTaken []string) (string, error) {
	slugIsTaken := func(str string, arr []string) bool {
		for _, i := range arr {
			if str == i {
				return true
			}
		}
		return false
	}

	slugBasename := Slugify(name)
	// if slug is unique without appending an integer, return it
	if !(slugIsTaken(slugBasename, slugsTaken)) {
		return slugBasename, nil
	}
	// max 1000 iterations trying to get unique slug
	// if we reach the end of 100 iterations, it means there are more than 100 things with the same
	// name in the database table
	i := 1
	for i < 1000 {
		slug := fmt.Sprintf("%s-%d", slugBasename, i)
		if !(slugIsTaken(slug, slugsTaken)) {
			return slug, nil
		}
		i++
	}
	return "", fmt.Errorf("reached max iteration 1000 without finding a unique slug")
}