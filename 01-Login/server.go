package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/tuxlife/auth0-golang-web-app/01-Login/app"
	"github.com/tuxlife/auth0-golang-web-app/01-Login/routes/callback"
	"github.com/tuxlife/auth0-golang-web-app/01-Login/routes/home"
	"github.com/tuxlife/auth0-golang-web-app/01-Login/routes/middlewares"
	"github.com/tuxlife/auth0-golang-web-app/01-Login/routes/user"
)

func init() {
	app.Init()
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
}
