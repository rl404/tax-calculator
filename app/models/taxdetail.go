package models

type TaxDetail struct {
	TaxCode 	int    		`json:"taxcode"`
	Name  		string    	`json:"name"`
	Refundable 	int    		`json:"refundable"`
}

var detailTax = map[int]TaxDetail{
	1: TaxDetail{
		TaxCode: 1,
		Name: "food & beverage",
		Refundable: 1,
	},
	2: TaxDetail{
		TaxCode: 2,
		Name: "tobacco",
		Refundable: 0,
	},
	3: TaxDetail{
		TaxCode: 3,
		Name: "entertainment",
		Refundable: 0,
	},
}