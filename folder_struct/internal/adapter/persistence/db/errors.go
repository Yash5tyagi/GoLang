package db

import "errors"

var (
	ErrDocumentNotFound = errors.New("document not found")
	ErrDuplicateEntry   = errors.New("duplicate entry")
)
