package util

import (
	"errors"
	"net/url"
)

type RedactRequest struct {
	URL string
}

// RedactQueryParam redacts a given url parameter returns the redacted *url.Error
//
// https://github.com/golang/go/issues/47442#issuecomment-888554396
func (r *RedactRequest) RedactQueryParam(queryParam string) error {
	if r == nil {
		return errors.New("nil RedactRequest")
	}

	u, uErr := url.Parse(r.URL)
	if uErr != nil {
		return errors.New("unable to parse URL string from urlErr")
	}
	q := u.Query()
	if p := q.Get(queryParam); p != "" {
		q.Set(queryParam, "REDACTED")
		u.RawQuery = q.Encode()
		r.URL = u.String()
	}

	return nil
}
