package auth

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func NewRegisterPage() *LoginPage {
	return &LoginPage{}
}

type RegisterPage struct {
	app.Compo
	username        string
	email           string
	password        string
	confirmPassword string
	updateAvailable bool
}

func (o *RegisterPage) Render() app.UI {
	return app.Div().Body(
		app.P().
			Class("sign").
			Text("Вход"),
		app.Form().
			Class("form1"),

		app.Input().
			Class("un").Type("email").
			Placeholder("Email").
			OnChange(func(ctx app.Context, e app.Event) {
				o.email = ctx.JSSrc().Get("value").String()
			}),

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
			OnInput(func(ctx app.Context, e app.Event) {
				o.password = ctx.JSSrc().Get("value").String()
			}),

		app.Input().
			Class("pass").
			Type("text").
			Placeholder("Confirm Password").
			OnInput(func(ctx app.Context, e app.Event) {
				o.confirmPassword = ctx.JSSrc().Get("value").String()
			}),

		app.If(o.updateAvailable,
			app.P(),
		),

		app.Button().
			Class("submit").
			Text("Войти"),

		app.P().
			Class("forgot").
			Text("Забыли пароль?"),
	)
}

func (o *RegisterPage) OnAppUpdate(ctx app.Context) {
	o.updateAvailable = ctx.AppUpdateAvailable() // Reports that an app update is available.
}

func (o *RegisterPage) isPasswordsEqual() bool {
	return o.password == o.confirmPassword
}
