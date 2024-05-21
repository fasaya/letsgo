package helpers

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

func GetIdFromURLParam(r *http.Request) (string, error) {
	// Split the URL path by '/' to get the individual path segments
	pathParts := strings.Split(r.URL.Path, "/")

	// The last segment should be the product ID
	id := pathParts[len(pathParts)-1]

	// check if id is number
	var re = regexp.MustCompile(`^[0-9]+$`)

	if re.MatchString(id) {
		// String contains only digits.
		return id, nil
	} else {
		// String does not contain only digits.

		return "", errors.New("ID is not a number")
	}

}
