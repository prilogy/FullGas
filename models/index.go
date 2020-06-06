package models

import (
	s "github.com/prilogy/FullGas/models/static"
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