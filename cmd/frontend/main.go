package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sorohimm/uacs-store-back/internal/service/frontend/admin/component/auth"
	"github.com/sorohimm/uacs-store-back/internal/service/frontend/admin/component/home"
	"github.com/sorohimm/uacs-store-back/internal/service/frontend/request"
	"log"
	"net/http"
)

func main() {
	requester := request.NewRequester("v1", "http://localhost:2104")

	app.Route("/", auth.NewLoginPage(requester))
	app.Route("/login", auth.NewLoginPage(requester))
	app.Route("/register", auth.NewRegisterPage(requester))
	app.Route("/admin", home.NewAdminHome())
	app.RunWhenOnBrowser()

	// Standard HTTP routing (server-side):
	http.Handle("/", &app.Handler{
		Name:  "UACS Archery Shop",
		Title: "UACS",
		Styles: []string{
			"/web/login.css",
			"/web/register.css",
			"/web/admin-dashboard.css",
		},
		LoadingLabel: "",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
