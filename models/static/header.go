package static

type HeaderData struct{
	Links		[]Link
	NavBar		[]Menu
	TopBar		[]Menu
}

type Link struct {
	Name	string
	Rel		string
	Href	string
}

type Menu struct {
	Name	string
	Link	string
	Rate	int
}

func Header() HeaderData {
	data := HeaderData{
		Links: []Link{
			{
				Name: "Главные стили",
				Rel: "stylesheet",
				Href: "/src/css/style.css",
			},
			{
				Name: "Шрифт логотипа",
				Rel: "stylesheet",
				Href: "https://fonts.googleapis.com/css2?family=Orbitron:wght@400;500;600;700;800;900&display=swap",
			},
			{
				Name: "Бутстрап",
				Rel: "stylesheet",
				Href: "/src/css/bootstrap-grid.min.css",
			},
		},
		NavBar: []Menu{
			{
				Name: "Главная",
				Link: "/",
				Rate: 1,
			},
			{
				Name: "Шины",
				Link: "/tires",
				Rate: 2,
			},
			{
				Name: "Колодки",
				Link: "/pads",
				Rate: 3,
			},
			{
				Name: "Цепи",
				Link: "/chains",
				Rate: 4,
			},
			{
				Name: "Звёзды",
				Link: "/stars",
				Rate: 5,
			},
		},
		TopBar: []Menu{
			{
				Name: "О нас",
				Link: "/",
				Rate: 1,
			},
			{
				Name: "FAQ",
				Link: "/",
				Rate: 2,
			},
			{
				Name: "Контакты",
				Link: "/",
				Rate: 3,
			},

		},
	}

	return data
}