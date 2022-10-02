package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sorohimm/uacs-store-back/internal/service/frontend/admin/component/auth"
	"log"
	"net/http"
)

func main() {
	// Go-app component routing (client-side):
	app.Route("/", &auth.LoginPage{})
	app.Route("/login", &auth.LoginPage{})
	app.Route("/register", &auth.RegisterPage{})
	app.RunWhenOnBrowser()

	// Standard HTTP routing (server-side):
	http.Handle("/", &app.Handler{
		Name:         "UACS Archery Shop",
		Title:        "UACS",
		Styles:       []string{},
		LoadingLabel: "",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
