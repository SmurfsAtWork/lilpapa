package actions

import "github.com/SmurfsAtWork/lilpapa/app"

type Actions struct {
	app         *app.App
	cache       Cache
	eventhub    EventHub
	blobstorage BlobStorage
	jwt         JwtManager[TokenPayload]
}

func New(
	app *app.App,
	cache Cache,
	eventhub EventHub,
	blobstorage BlobStorage,
	jwt JwtManager[TokenPayload],
) *Actions {
	return &Actions{
		app:         app,
		cache:       cache,
		eventhub:    eventhub,
		blobstorage: blobstorage,
		jwt:         jwt,
	}
}
