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
		NewSidebarList(),
	).Class("home-admin-sidebar")
}

func NewSidebarList() *SidebarNav {
	return &SidebarNav{}
}

type SidebarNav struct {
	app.Compo
}

func (o *SidebarNav) Render() app.UI {
	return app.Nav().
		Class("home-admin-nav").
		Body(
			NewSidebarSection("Дашборд"),
			NewSidebarSection("Товары"),
			NewSidebarSection("Категории"),
			NewSidebarSection("Бренды"),
			NewSidebarSection("Заказы"),
		)
}

func NewSidebarSection(text string) *SidebarSection {
	return &SidebarSection{
		text: text,
	}
}

type SidebarSection struct {
	app.Compo
	text string
}

func (o *SidebarSection) Render() app.UI {
	return app.Div().
		Class("sidebar-text-section-container").
		Body(
			app.Span().
				Class("sidebar-text-section-text").
				Body(
					app.Span().Text(o.text),
				),
			app.Raw(`<svg viewBox="0 0 1024 1024" class="sidebar-text-section-icon">
        						<path d="M366 708l196-196-196-196 60-60 256 256-256 256z"></path>
     					 </svg>`,
			),
		)
}
