package models

import (
	s "../models/static"
)

type PadsData struct{
	Header		s.HeaderData
	Pads		[]PadsProduct
	PageMark	string
	MarkList	[]PadsMark
	Page	 	Pagination
}

type PadsMark struct{
	Id		int
	Name	string
}

type Pagination struct {
	Start	int
	End		int
	AllData int
	Current int
	Link	string
}

type PadsProduct struct {
	Id			int
	Mark		int
	MarkName	string
	Model		string
	Years		string
	Img			int
	Price		int
}

func Pads() PadsData{
	data := PadsData{
		Header: s.Header(),
		MarkList: []PadsMark{
			{Id: 1, Name: "KTM"},
			{Id: 2, Name: "Husqvarna"},
			{Id: 3, Name: "Husaberg"},
			{Id: 4, Name: "BMW"},
			{Id: 5, Name: "Gas-Gas"},
			{Id: 6, Name: "Sherco"},
			{Id: 7, Name: "TM Racing"},
			{Id: 8, Name: "YAMAHA"},
			{Id: 9, Name: "Honda"},
			{Id: 10, Name: "SUZUZKI"},
			{Id: 11, Name: "Kawasaki"},
		},
	}

	return data
}

