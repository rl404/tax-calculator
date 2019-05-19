package models

type TaxDetail struct {
	TaxCode 	int    		`json:"taxcode"`
	Name  		string    	`json:"name"`
	Refundable 	int    		`json:"refundable"`
}

// Static detail since it is not saved to db
var detailTax = map[int]TaxDetail{
	1: TaxDetail{
		TaxCode: 1,
		Name: "Food & Beverage",
		Refundable: 1,
	},
	2: TaxDetail{
		TaxCode: 2,
		Name: "Tobacco",
		Refundable: 0,
	},
	3: TaxDetail{
		TaxCode: 3,
		Name: "Entertainment",
		Refundable: 0,
	},
}

func (t TaxDetail) GetDetailList() map[int]TaxDetail {
	return detailTax
}