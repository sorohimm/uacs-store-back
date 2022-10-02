package home

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func NewSidebar() *Sidebar {
	return &Sidebar{}
}

type Sidebar struct {
	app.Compo
}

func (o *Sidebar) Render() app.UI {
	return app.Div().Body(
		app.Ul().Body(
			app.Li().Body(
				app.A().Body(
					app.Span().
						Class("item").
						Text("Dashboard"),
				).Class("active").Href("#"),
			),
			app.Li().Body(
				app.A().Body(
					app.Span().
						Class("item").
						Text("Products"),
				).Class("active").Href("#"),
			),
			app.Li().Body(
				app.A().Body(
					app.Span().
						Class("item").
						Text("Brands"),
				).Class("active").Href("#"),
			),
			app.Li().Body(
				app.A().Body(
					app.Span().
						Class("item").
						Text("Categories"),
				).Class("active").Href("#"),
			),
			app.Li().Body(
				app.A().Body(
					app.Span().
						Class("item").
						Text("Orders"),
				).Class("active").Href("#"),
			),
		),
	).Class("admin-sidebar")
}
