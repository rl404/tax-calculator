package models

type Tax struct {
	Id			int 		`json:"-" db:"id"`
	BillId 		int 		`json:"-" db:"bill_id"`

	Name  		string    	`json:"name" db:"name"`
	TaxCode 	int    		`json:"taxcode" db:"tax_code"`
	Price 		float64		`json:"price" db:"price"`

	TaxType 	string 		`json:"type" db:"-"`
	Refundable 	string 		`json:"refundable" db:"-"`
	Tax 		float64 	`json:"tax" db:"-"`
	Amount		float64 	`json:"amount" db:"-"`
}

func (c *Tax) GetAllDetail() {
	c.TaxType = c.GetTaxType()
	c.Refundable = c.GetRefundable()
	c.Tax = c.GetTax()
	c.Amount = c.GetAmount()
}

func (c Tax) GetTax() float64 {
	p, t := c.Price, c.TaxCode

	// food & beverage, 10% of price
	if t == 1 {
		return p * 10.0/100.0
	}

	// tobacco, 10 + (2% of p)
	if t == 2 {
		return 10.0 + (p * 2.0/100.0)
	}

	// entertainment, free if 0 < p < 100 else 1% of (p - 100)
	if t == 3 {
		if p > 0 && p < 100 {
			return p
		}

		if p >= 100 {
			return (p - 100.0) * 1.0/100.0
		}
	}

	return 0
}

func (c Tax) GetTaxType() string {
	_, exists := detailTax[c.TaxCode]
	if exists {
		return detailTax[c.TaxCode].Name
	}
	return ""
}

func (c Tax) GetRefundable() string {
	_, exists := detailTax[c.TaxCode]
	if exists {
		if detailTax[c.TaxCode].Refundable == 1 {
			return "yes"
		}
		return "no"
	}
	return ""
}

func (c Tax) GetAmount() float64 {
	return c.Price + c.Tax
}