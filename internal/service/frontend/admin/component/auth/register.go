package auth

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sorohimm/uacs-store-back/internal/service/frontend/request"
)

func NewRegisterPage(requester *request.Requester) *RegisterPage {
	return &RegisterPage{
		requester: requester,
	}
}

type RegisterPage struct {
	app.Compo
	username        string
	email           string
	password        string
	confirmPassword string
	requester       *request.Requester
}

func (o *RegisterPage) Render() app.UI {
	return app.Div().Body(
		app.H1().
			Class("register-text").
			Text("Регистрация"),

		app.Input().
			Class("register-textinput input").
			Type("text").
			Placeholder("Username").
			OnInput(func(ctx app.Context, e app.Event) {
				o.username = ctx.JSSrc().Get("value").String()
				app.Log("username is: ", o.username)
			}),

		app.Input().
			Class("register-textinput1 input").
			Type("email").
			Placeholder("Email").
			OnInput(func(ctx app.Context, e app.Event) {
				o.email = ctx.JSSrc().Get("value").String()
				app.Log("email is: ", o.email)
			}),

		app.Input().
			Class("register-textinput2 input").
			Type("password").
			Placeholder("Password").
			OnInput(func(ctx app.Context, e app.Event) {
				o.password = ctx.JSSrc().Get("value").String()
				app.Log("password is: ", o.password)
			}),

		app.Input().
			Class("register-textinput3 input").
			Type("password").
			Placeholder("Confirm Password").
			OnInput(func(ctx app.Context, e app.Event) {
				o.confirmPassword = ctx.JSSrc().Get("value").String()
				app.Log("confirmPassword is: ", o.confirmPassword)
			}),

		app.Button().
			Class("register-button button").
			Text("Зарегистрироваться").
			OnClick(o.onClickRegisterButton),
	).Class("register-container")
}

func (o *RegisterPage) onClickRegisterButton(ctx app.Context, e app.Event) {
	if o.isPasswordsEqual() == false {
		return
	}
	if o.isPasswordsEqual() == false {
		return
	}

	req := request.NewRegisterRequest().
		SetUsername(o.username).
		SetPassword(o.password).
		SetEmail(o.email).
		SetRole("USER")

	err := o.requester.Register(req)
	if err == nil {
		ctx.Navigate("/login")
	}
	app.Log(err)
}

func (o *RegisterPage) isPasswordsEqual() bool {
	return o.password == o.confirmPassword
}

func (o *RegisterPage) isAllFieldsFill() bool {
	if o.username == "" || o.email == "" || o.password == "" || o.confirmPassword == "" {
		return false
	}
	return true
}
