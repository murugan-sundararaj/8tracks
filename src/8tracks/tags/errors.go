package tags

import "errors"

var (
	ErrInvalidTag   = errors.New("one or more invalid tag found")
	ErrTagNameExist = errors.New("tag name already exist")
)
