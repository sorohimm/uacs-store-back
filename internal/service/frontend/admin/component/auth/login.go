package auth

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func NewLoginPage() *LoginPage {
	return &LoginPage{}
}

type LoginPage struct {
	app.Compo
	username string
	password string
}

func (o *LoginPage) Render() app.UI {
	return app.Div().Body(
		app.P().
			Class("sign").
			Text("Вход"),
		app.Form().
			Class("form1"),
		app.Input().
			Class("un").Type("text").
			Placeholder("Username").
			OnChange(func(ctx app.Context, e app.Event) {
				o.username = ctx.JSSrc().Get("value").String()
			}),
		app.Input().
			Class("pass").
			Type("text").
			Placeholder("Password").
			OnChange(func(ctx app.Context, e app.Event) {
				o.password = ctx.JSSrc().Get("value").String()
			}),
		app.Button().
			Class("submit").
			Text("Войти"),
		app.P().
			Class("forgot").
			Text("Забыли пароль?"),
	)
}
