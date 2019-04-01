package main

import (
	"github.com/go-errors/errors"
)

func crash() error {
	return errors.New(errors.Errorf("oh dear"))
}
