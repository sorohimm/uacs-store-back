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
		/*<ul>
		  <li>lorem ipsum dolor</li>
		  <li>sit amet</li>
		  <li>foo</li>
		  <li>bar</li>
		</ul>*/
		app.Div().Body(
			app.Ul().Body(
				app.Li().Text("lorem ipsum dolor"),
				app.Li().Text("sit amet"),
				app.Li().Text("foo"),
				app.Li().Text("bar"),
			),
		),
	).Class("wrapper")
}
