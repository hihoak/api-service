package custom_errors

import "errors"

var (
	ErrArtistNotFound = errors.New("artist doesn't found")
)