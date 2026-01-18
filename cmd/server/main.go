package main

import (
	"net/http"
	"regexp"

	"github.com/SmurfsAtWork/lilpapa/actions"
	"github.com/SmurfsAtWork/lilpapa/app"
	"github.com/SmurfsAtWork/lilpapa/config"
	"github.com/SmurfsAtWork/lilpapa/evy"
	"github.com/SmurfsAtWork/lilpapa/handlers/apis"
	"github.com/SmurfsAtWork/lilpapa/handlers/middlewares/auth"
	"github.com/SmurfsAtWork/lilpapa/handlers/middlewares/contenttype"
	"github.com/SmurfsAtWork/lilpapa/handlers/middlewares/logger"
	"github.com/SmurfsAtWork/lilpapa/handlers/middlewares/smurfauth"
	"github.com/SmurfsAtWork/lilpapa/jwt"
	"github.com/SmurfsAtWork/lilpapa/log"
	"github.com/SmurfsAtWork/lilpapa/memcache"
	"github.com/SmurfsAtWork/lilpapa/sqlite"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
)

var (
	minifyer            *minify.M
	usecases            *actions.Actions
	authMiddleware      *auth.Middleware
	smurfAuthMiddleware *smurfauth.Middleware
)

func init() {
	sqlite3Repo, err := sqlite.New()
	if err != nil {
		log.Fatalln(err)
	}

	app := app.New(sqlite3Repo)
	cache := memcache.New()
	eventhub := evy.New(sqlite3Repo)
	jwt := jwt.New[actions.TokenPayload]()

	usecases = actions.New(
		app,
		cache,
		eventhub,
		nil,
		jwt,
	)

	minifyer = minify.New()
	minifyer.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)

	authMiddleware = auth.New(usecases)
	smurfAuthMiddleware = smurfauth.New(usecases)
}

func main() {
	// MIGRATOR
	err := sqlite.Migrate()
	if err != nil {
		log.Fatalln(err)
	}

	// EVENTS
	go fetchAndExecuteEventsAsync()

	// CDN

	// HTTP
	adminLoginApi := apis.NewAdminLoginApi(usecases)
	meApi := apis.NewMeApi(usecases)
	smurfLoginApi := apis.NewSmurfLoginApi(usecases)
	smurfMeApi := apis.NewSmurfMeApi(usecases)

	v1ApisHandler := http.NewServeMux()
	v1ApisHandler.HandleFunc("POST /login", adminLoginApi.HandleLogin)
	v1ApisHandler.HandleFunc("POST /login/smurf", smurfLoginApi.HandleLogin)

	v1ApisHandler.HandleFunc("GET /me", authMiddleware.AuthApi(meApi.HandleAuthCheck))
	v1ApisHandler.HandleFunc("GET /me/logout", authMiddleware.AuthApi(meApi.HandleLogout))
	v1ApisHandler.HandleFunc("GET /me/smurf", smurfAuthMiddleware.AuthApi(smurfMeApi.HandleAuthCheck))
	v1ApisHandler.HandleFunc("GET /me/smurf/logout", smurfAuthMiddleware.AuthApi(smurfMeApi.HandleLogout))

	// v1ApisHandler.HandleFunc("POST /smurf", nil)

	applicationHandler := http.NewServeMux()
	applicationHandler.Handle("/v1/", http.StripPrefix("/v1", contenttype.Json(v1ApisHandler)))

	log.Info("Starting http server at port " + config.Env().Port)
	switch config.Env().GoEnv {
	case config.GoEnvBeta, config.GoEnvDev, config.GoEnvTest:
		log.Fatalln(http.ListenAndServe(":"+config.Env().Port, logger.Handler(applicationHandler)))
	case config.GoEnvProd:
		log.Fatalln(http.ListenAndServe(":"+config.Env().Port, minifyer.Middleware(applicationHandler)))
	}
}
