package controllers

import (
	"github.com/revel/revel"
	"github.com/rl404/tax-calculator/app/models"
	"strconv"
	"time"
)

type Api struct {
	*revel.Controller
}

func (a Api) CreateBill() revel.Result {
	var taxModel []models.Tax
	a.Params.BindJSON(&taxModel)

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
	bill.CreatedDate = time.Now().Local().Unix()
	Dbm.Update(&bill)

	return a.RenderJSON(bill)
}

func (a Api) GetBill() revel.Result {
	billid := a.Params.Query.Get("bill")

	intBillId,_ := strconv.Atoi(billid)
	billModel := a.GetBillDetail(intBillId)

	return a.RenderJSON(billModel)
}

func (a Api) GetBillDetail(billid int) models.Bill {
	var billModel models.Bill
	var taxList []models.Tax

	_, err := Dbm.Select(&taxList, `SELECT name, tax_code, price FROM tax WHERE bill_id =?`, billid)

	if err != nil {
		panic(err)
	}

	taxList = a.GetTaxDetail(taxList)

	_ = Dbm.SelectOne(&billModel, `SELECT bill_id, price_total, tax_total, grand_total, created_date FROM bill WHERE bill_id =?`, billid)

	billModel.Detail = taxList

	return billModel
}

func (a Api) GetTaxDetail(taxList []models.Tax) []models.Tax {
	for i,_ := range taxList {
		taxList[i].GetAllDetail()
	}
	return taxList
}

func (a Api) Dummy() []models.Tax {
	var taxModel []models.Tax
	taxModel = append(taxModel, models.Tax{Name:"Lucky Stretch",TaxCode:2,Price:1000})
	taxModel = append(taxModel, models.Tax{Name:"Big Mac",TaxCode:1,Price:1000})
	taxModel = append(taxModel, models.Tax{Name:"Movie",TaxCode:3,Price:150})
	return taxModel
}