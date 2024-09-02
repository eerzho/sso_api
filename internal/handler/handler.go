package handler

import (
	"net/http"
	"sso/internal/app"
	"sso/internal/handler/mwr"
	v1 "sso/internal/handler/v1"
)

func New(
	app *app.App,
	userSrvc v1.UserSrvc,
) http.Handler {
	mux := http.NewServeMux()
	reqIDMwr := mwr.NewRequestId("X-Request-ID")
	reqLgMwr := mwr.NewRequestLogger(app.Lg)

	v1.New(mux, app, "/api/v1", userSrvc)

	return reqIDMwr.Mwr(reqLgMwr.Mwr(mux))
}