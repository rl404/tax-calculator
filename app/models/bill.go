package models

type Bill struct {
	Id 			int 		`json:"-" db:"id"`
	BillId 		int 		`json:"billid" db:"bill_id"`
	Detail 		[]Tax 		`json:"detail" db:"-"`
	PriceTotal	float64 	`json:"pricetotal" db:"price_total"`
	TaxTotal	float64 	`json:"taxtotal" db:"tax_total"`
	GrandTotal 	float64 	`json:"grandtotal" db:"grand_total"`
	CreatedDate	int64 		`json:"createddate" db:"created_date"`
}