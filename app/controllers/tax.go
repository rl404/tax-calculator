package controllers

import (
	"github.com/revel/revel"
	"tax-calculator/app/models"
)

type Tax struct {
	*revel.Controller
}

func (c Tax) Index() revel.Result {

	taxList, err := Dbm.Select(models.Tax{},
		`SELECT name, tax_code, price FROM tax`)
	if err != nil {
		panic(err)
	}

	return c.RenderJSON(taxList)
}

func (c Tax) Calculate() revel.Result {
	var taxModel []models.Tax
	c.Params.BindJSON(&taxModel)

	var bill models.Bill
	Dbm.Insert(&bill)
	bill.BillId = bill.Id

	for i,_ := range taxModel {
		taxModel[i].GetAllDetail()
		taxModel[i].BillId = bill.BillId
		Dbm.Insert(&taxModel[i])

		bill.PriceTotal += taxModel[i].Price
		bill.TaxTotal += taxModel[i].Tax
		bill.GrandTotal += taxModel[i].Amount
	}

	bill.Detail = taxModel
	Dbm.Update(&bill)

	return c.RenderJSON(bill)
}

func (c Tax) Dummy() []models.Tax {
	var taxModel []models.Tax
	taxModel = append(taxModel, models.Tax{Name:"Lucky Stretch",TaxCode:2,Price:1000})
	taxModel = append(taxModel, models.Tax{Name:"Big Mac",TaxCode:1,Price:1000})
	taxModel = append(taxModel, models.Tax{Name:"Movie",TaxCode:3,Price:150})
	return taxModel
}

func (c Tax) TestPost() revel.Result {
	return c.Render()
}