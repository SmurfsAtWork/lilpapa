package memcache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SmurfsAtWork/lilpapa/app/models"
)

const (
	userSessionTokenTtlDays  = 7
	smurfSessionTokenTtlDays = 30
)

type Cache struct {
	client *memoryCache
}

func New() *Cache {
	return &Cache{
		client: newMemoryCache(),
	}
}

func userTokenKey(sessionToken string) string {
	return fmt.Sprintf("user-session-token:%s", sessionToken)
}

func (c *Cache) SetAuthenticatedUser(sessionToken string, user models.User) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.client.Set(userTokenKey(sessionToken), string(userJson), userSessionTokenTtlDays*time.Hour*24)
}

func (c *Cache) GetAuthenticatedUser(sessionToken string) (models.User, error) {
	value, err := c.client.Get(userTokenKey(sessionToken))
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *Cache) InvalidateAuthenticatedUser(sessionToken string) error {
	err := c.client.Del(userTokenKey(sessionToken))
	if err != nil {
		return err
	}

	return nil
}

func smurfTokenKey(sessionToken string) string {
	return fmt.Sprintf("smurf-session-token:%s", sessionToken)
}

func (c *Cache) SetAuthenticatedSmurf(sessionToken string, smurf models.Smurf) error {
	smurfJson, err := json.Marshal(smurf)
	if err != nil {
		return err
	}

	return c.client.Set(smurfTokenKey(sessionToken), string(smurfJson), smurfSessionTokenTtlDays*time.Hour*24)
}

func (c *Cache) GetAuthenticatedSmurf(sessionToken string) (models.Smurf, error) {
	value, err := c.client.Get(smurfTokenKey(sessionToken))
	if err != nil {
		return models.Smurf{}, err
	}

	var smurf models.Smurf
	err = json.Unmarshal([]byte(value), &smurf)
	if err != nil {
		return models.Smurf{}, err
	}

	return smurf, nil
}

func (c *Cache) InvalidateAuthenticatedSmurf(sessionToken string) error {
	err := c.client.Del(smurfTokenKey(sessionToken))
	if err != nil {
		return err
	}

	return nil
}
