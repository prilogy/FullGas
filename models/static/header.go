package static

type HeaderData struct{
	Links		[]Link
}

type Link struct {
	Name	string
	Rel		string
	Href	string
}

func Header() HeaderData {
	data := HeaderData{
		Links: []Link{
			{
				Name: "Главные стили",
				Rel: "stylesheet",
				Href: "src/css/style.css",
			},
		},
	}

	return data
}