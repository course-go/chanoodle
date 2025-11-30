package id

import "errors"

var ErrNoSuchEntity = errors.New("no entity with such ID")

type ID int
