package models

import (
	s "../models/static"
)

type IndexData struct{
	Header		s.HeaderData
}

func Index() IndexData{
	data := IndexData{
		Header: s.Header(),
	}

	return data
}