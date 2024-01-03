package repository

import "time"

type InvoiceProjection struct {
	Id            string
	CreatedAt     time.Time
	BuyAt         time.Time
	PayIn         time.Time
	Description   string
	Value         float64
	IsAlreadyDone bool
	UserId        string
	Category      InvoiceCategory
	PaymentType   PaymentType
}

type Invoice struct {
	Id                  string
	CreatedAt           time.Time
	BuyAt               time.Time
	PayAt               time.Time
	Description         string
	Value               float64
	InvoiceProjectionId string
	UserId              string
	Category            InvoiceCategory
	PaymentType         PaymentType
}

type InvoiceCategory struct {
	Id       uint
	Category string
}

type PaymentType struct {
	Id   uint
	Type string
}

type QueryParams struct {
	userId string
	month  uint
	year   uint
	limit  uint
	offset uint
}
