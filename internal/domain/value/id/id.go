package id

import "errors"

var ErrNoSuchEntity = errors.New("no such entity")

type ID int
