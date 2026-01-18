package actions

import "github.com/SmurfsAtWork/lilpapa/app/models"

type Cache interface {
	SetAuthenticatedUser(sessionToken string, user models.User) error
	GetAuthenticatedUser(sessionToken string) (models.User, error)
	InvalidateAuthenticatedUser(sessionToken string) error

	SetAuthenticatedSmurf(sessionToken string, smurf models.Smurf) error
	GetAuthenticatedSmurf(sessionToken string) (models.Smurf, error)
	InvalidateAuthenticatedSmurf(sessionToken string) error
}
