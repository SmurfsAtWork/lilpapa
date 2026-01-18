package sqlite

import (
	"errors"

	"gorm.io/gorm"
)

type ErrRecordNotFound struct{}

func (e ErrRecordNotFound) Error() string {
	return "record-not-found"
}

func (e ErrRecordNotFound) ClientStatusCode() int {
	return 404
}

func (e ErrRecordNotFound) ExtraData() map[string]any {
	return nil
}

func (e ErrRecordNotFound) ExposeToClients() bool {
	return false
}

type ErrRecordExists struct{}

func (e ErrRecordExists) Error() string {
	return "record-exists"
}

func (e ErrRecordExists) ClientStatusCode() int {
	return 409
}

func (e ErrRecordExists) ExtraData() map[string]any {
	return nil
}

func (e ErrRecordExists) ExposeToClients() bool {
	return false
}

func tryWrapDbError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return &ErrRecordNotFound{}
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return &ErrRecordExists{}
	}

	return err
}
