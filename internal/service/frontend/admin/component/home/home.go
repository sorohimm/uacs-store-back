package home

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func NewAdminHome() *AdminHome {
	return &AdminHome{}
}

type AdminHome struct {
	app.Compo
}

func (o *AdminHome) Render() app.UI {
	return app.Div().Body(
		NewSidebar(),
	).Class("wrapper")
}
