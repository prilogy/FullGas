package models

import (
	s "FullGas/models/static"
)

type TiresData struct{
	Header		s.HeaderData
	TiresInside	TiresInside

}

type TiresInside struct {
	Cub				[]int
	RadiusInside	map[int][]int
	Type			[]TiresType
	Radius			[]int
	Spike			[]int
}

type TiresType struct{
	Id		int
	Name	string
}

func Tires() TiresData{
	data := TiresData{
		Header: s.Header(),
		TiresInside: TiresInside{
			Type: []TiresType{
				{
					Id: 1,
					Name: "Комплект",
				},{
					Id: 2,
					Name: "Передняя",
				},{
					Id: 3,
					Name: "Задняя",
				},
			},
		},
	}

	return data
}