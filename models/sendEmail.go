package models

import (
	s "FullGas/models/static"
)

type SendEmailData struct{
	Header		s.HeaderData
	OrderData	SendOrder
}

type SendOrder struct{
	Id			string
	Product		string
	ProductId	int
}

func SendEmail() SendEmailData{
	data := SendEmailData{
		Header: s.Header(),
	}

	return data
}

