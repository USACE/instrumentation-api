package dbutils

import (
	"errors"
	"fmt"
	"log"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
)

// Slugify removes spaces and converts to lower case
func Slugify(str string) string {
	slug := slug.Make(str)
	return slug
}

func slugIsUnique(DB *sqlx.DB, slug string, table string, column string) bool {

	sql := fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE %s = $1", column, table, column)

	log.Printf(sql)
	log.Printf(column)

	var count int
	err := DB.Get(&count, sql, slug)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("Slugs with this name: %s; %d", slug, count)
	if count != 0 {
		return false
	}
	return true
}

// NextUniqueSlug returns the next unique slug available based on
func NextUniqueSlug(DB *sqlx.DB, str string, table string, column string) (string, error) {

	slugBasename := Slugify(str)
	// if slug is unique without appending an integer
	if slugIsUnique(DB, slugBasename, table, column) {
		return slugBasename, nil
	}
	// max 10 iterations trying to get unique slug
	i := 1
	for i < 10 {
		slug := fmt.Sprintf("%s-%d", slugBasename, i)
		if slugIsUnique(DB, slug, table, column) {
			return slug, nil
		}
		i++
	}
	return "", errors.New("reached max iteration %i without finding a unique slug")
}
