package models

import (
	s "golex/models/static"
)

type ViewData struct{
	Header		s.HeaderData
}

type ProductData struct {
	Name	string
	Header	string
	Blocks	[]Block
}

type Block struct {
	Img		string
	Name 	string
	Desc	string
}

func Index() ViewData{
	data := ViewData{
		Header: s.Header(),
	}

	return data
}