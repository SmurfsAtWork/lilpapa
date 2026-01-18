package apis

import (
	"context"

	"github.com/SmurfsAtWork/lilpapa/actions"
	"github.com/SmurfsAtWork/lilpapa/app/models"
	"github.com/SmurfsAtWork/lilpapa/handlers/middlewares/auth"
)

func parseContext(ctx context.Context) (actions.ActionContext, error) {
	user, userCorrect := ctx.Value(auth.UserKey).(models.User)
	if !userCorrect {
		return actions.ActionContext{}, &ErrUnauthorized{}
	}

	return actions.ActionContext{
		User: user,
	}, nil
}
