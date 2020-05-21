package models

import (
	s "../models/static"
)

type ChainsData struct{
	Header		s.HeaderData
	Chains		[]ChainsProduct
	PageMark	string
	UsageList	[]ChainsUsage
	Page	 	Pagination
}

type ChainsUsage struct{
	Id		int
	Name	string
}

type ChainsProduct struct {
	Id			int
	Label		string
	Model		string
	Mark		int
	MarkName	string
	Description string
	Price		int
}

func Chains() ChainsData{
	data := ChainsData{
		Header: s.Header(),
		UsageList: []ChainsUsage{
			{Id: 1, Name: "Motocross"},
			{Id: 2, Name: "Road/offroad"},
			{Id: 3, Name: "Street"},
		},
	}

	return data
}

