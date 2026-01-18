package app

import (
	"fmt"
	"net/http"
	"strings"
)

// PapaError is implemented for every error around here :)
type PapaError interface {
	error
	// ClientStatusCode the HTTP status for clients.
	ClientStatusCode() int
	// ExtraData any data that will be helpful for clients for better UX context.
	ExtraData() map[string]any
	// ExposeToClients reports whether to expose this error to clients or not.
	ExposeToClients() bool
}

type ErrNotFound struct {
	ResourceName string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s-not-found", strings.ToLower(e.ResourceName))
}

func (e ErrNotFound) ClientStatusCode() int {
	return http.StatusNotFound
}

func (e ErrNotFound) ExtraData() map[string]any {
	return nil
}

func (e ErrNotFound) ExposeToClients() bool {
	return true
}

type ErrExists struct {
	ResourceName string
}

func (e ErrExists) Error() string {
	return fmt.Sprintf("%s-exists", strings.ToLower(e.ResourceName))
}

func (e ErrExists) ClientStatusCode() int {
	return http.StatusConflict
}

func (e ErrExists) ExtraData() map[string]any {
	return nil
}

func (e ErrExists) ExposeToClients() bool {
	return true
}

type ErrExpiredVerificationCode struct{}

func (e ErrExpiredVerificationCode) Error() string {
	return "verification-code-expired"
}

func (e ErrExpiredVerificationCode) ClientStatusCode() int {
	return http.StatusUnauthorized
}

func (e ErrExpiredVerificationCode) ExtraData() map[string]any {
	return nil
}

func (e ErrExpiredVerificationCode) ExposeToClients() bool {
	return true
}

type ErrInvalidVerificationToken struct{}

func (e ErrInvalidVerificationToken) Error() string {
	return "invalid-verification-code"
}

func (e ErrInvalidVerificationToken) ClientStatusCode() int {
	return http.StatusBadRequest
}

func (e ErrInvalidVerificationToken) ExtraData() map[string]any {
	return nil
}

func (e ErrInvalidVerificationToken) ExposeToClients() bool {
	return true
}
