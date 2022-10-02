package auth

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sorohimm/uacs-store-back/internal/service/frontend/request"
)

func NewLoginPage(requester *request.Requester) *LoginPage {
	return &LoginPage{
		requester: requester,
	}
}

type LoginPage struct {
	app.Compo
	username  string
	password  string
	requester *request.Requester
}

func (o *LoginPage) Render() app.UI {
	return app.Div().Body(
		app.H1().
			Class("login-text").
			Text("Вход"),

		app.Input().
			Class("login-username-input").
			Type("text").
			Placeholder("Username").
			OnChange(func(ctx app.Context, e app.Event) {
				o.username = ctx.JSSrc().Get("value").String()
			}),

		app.Input().
			Class("login-password-input").
			Type("password").
			Placeholder("Password").
			OnChange(func(ctx app.Context, e app.Event) {
				o.password = ctx.JSSrc().Get("value").String()
			}),

		app.A().Class("login-forgot-password-link").Href("").Body(
			app.Span().Text("Забыли пароль?"),
		),

		app.Button().
			Class("login-signin-button").
			Text("Войти").
			OnClick(o.onClickLoginButton),
	).Class("login-container")
}

func (o *LoginPage) onClickLoginButton(ctx app.Context, _ app.Event) {
	if o.password != "" && o.username != "" {
		req := request.NewLoginRequest().SetUsername(o.username).SetPassword(o.password)
		if _, err := o.requester.Login(req); err == nil {
			ctx.Navigate("/admin")
		}
	}
}
