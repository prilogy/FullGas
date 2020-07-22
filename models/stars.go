package models

import (
	s "FullGas/models/static"
)

type StarsData struct{
	Header		s.HeaderData
	Pads		[]StarsProduct
	PageMark	string
	MarkList	[]StarsMark
	Page	 	Pagination
}

type StarsMark struct{
	Id		int
	Name	string
}

type StarsProduct struct {
	Id			int
	Brand		string
	Mark		string
	Model		string
	Years		string
	Side		int
	Img			int
	Price		int
}

func Stars() StarsData{
	data := StarsData{
		Header: s.Header(),
		MarkList: []StarsMark{
			{Id: 1, Name: "Honda"},
			{Id: 2, Name: "Kawasaki"},
			{Id: 3, Name: "KTM"},
			{Id: 4, Name: "Suzuki"},
			{Id: 5, Name: "Husqvarna"},
			{Id: 6, Name: "YAMAHA"},
		},
	}

	return data
}

