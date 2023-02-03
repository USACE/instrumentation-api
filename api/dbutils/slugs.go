package dbutils

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
)

// Slugify removes spaces and converts to lower case
func Slugify(str string) string {
	slug := slug.Make(str)
	return slug
}

// NextUniqueSlug returns the next unique slug available based on
func NextUniqueSlug(str string, usedSlugs []string) (string, error) {

	slugIsTaken := func(str string, arr []string) bool {
		for _, i := range arr {
			if str == i {
				return true
			}
		}
		return false
	}

	slugBasename := Slugify(str)
	// if slug is unique without appending an integer, return it
	if !(slugIsTaken(slugBasename, usedSlugs)) {
		return slugBasename, nil
	}
	// max 1000 iterations trying to get unique slug
	// if we reach the end of 100 iterations, it means there are more than 100 things with the same
	// name in the database table
	i := 1
	for i < 1000 {
		slug := fmt.Sprintf("%s-%d", slugBasename, i)
		if !(slugIsTaken(slug, usedSlugs)) {
			return slug, nil
		}
		i++
	}
	return "", errors.New("reached max iteration %i without finding a unique slug")
}

func ListSlugs(db *sqlx.DB, tableName string) ([]string, error) {
	slugs := make([]string, 0)
	if err := db.Select(&slugs, "SELECT slug FROM "+tableName); err != nil {
		return make([]string, 0), err
	}
	return slugs, nil
}

// CreateUniqueSlug creates a unique slug given a name and tableName
func CreateUniqueSlug(db *sqlx.DB, name string, tableName string) (string, error) {
	slugsTaken, err := ListSlugs(db, tableName)
	if err != nil {
		return "", err
	}
	s, err := NextUniqueSlug(name, slugsTaken)
	if err != nil {
		return "", err
	}
	return s, nil
}
