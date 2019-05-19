package controllers

import (
	"github.com/revel/revel"
	"strconv"
	"time"
	"math"

	"github.com/rl404/tax-calculator/app/models"
	"github.com/rl404/tax-calculator/app/helpers"
)

type Api struct {
	*revel.Controller
}

func (a Api) CreateBill() revel.Result {
	var taxModel []models.Tax
	a.Params.BindJSON(&taxModel)

	// Insert to db
	var bill models.Bill
	err := Dbm.Insert(&bill)
	if err != nil {
		response := helpers.ToRest(500, "Failed insert bill to database", nil)
		return a.RenderJSON(response)
	}

	// Get detail, insert to db, count total
	for i,_ := range taxModel {
		taxModel[i].BillId = bill.Id
		taxModel[i].GetAllDetail()
		err = Dbm.Insert(&taxModel[i])
		if err != nil {
			response := helpers.ToRest(500, "Failed insert tax to database", nil)
			return a.RenderJSON(response)
		}

		bill.PriceTotal += taxModel[i].Price
		bill.TaxTotal += taxModel[i].Tax
		bill.GrandTotal += taxModel[i].Amount
	}

	// Update bill id
	bill.BillId = bill.Id
	bill.Detail = taxModel
	bill.PriceTotal = math.Round(bill.PriceTotal*100)/100
	bill.TaxTotal = math.Round(bill.TaxTotal*100)/100
	bill.GrandTotal = math.Round(bill.GrandTotal*100)/100
	bill.CreatedDate = time.Now().Local().Unix()

	_, err = Dbm.Update(&bill)
	if err != nil {
		response := helpers.ToRest(500, "Failed update bill to database", nil)
		return a.RenderJSON(response)
	}

	response := helpers.ToRest(200, "", bill)
	return a.RenderJSON(response)
}

func (a Api) GetBill() revel.Result {
	billid := a.Params.Query.Get("bill")

	intBillId,_ := strconv.Atoi(billid)
	billModel, err, errCode := a.GetBillDetail(intBillId)

	if err != "" {
		response := helpers.ToRest(errCode, err, nil)
		return a.RenderJSON(response)
	}

	response := helpers.ToRest(200, "", billModel)
	return a.RenderJSON(response)
}

func (a Api) GetBillDetail(billid int) (models.Bill, string, int) {
	var billModel models.Bill
	var taxList []models.Tax

	_, err := Dbm.Select(&taxList, `SELECT name, tax_code, price FROM tax WHERE bill_id =?`, billid)

	if err != nil {
		return models.Bill{}, "Failed select tax from database", 500
	}

	if len(taxList) == 0 {
		return models.Bill{}, "Bill Not Found", 404
	}

	taxList = a.GetTaxDetail(taxList)

	err = Dbm.SelectOne(&billModel, `SELECT bill_id, price_total, tax_total, grand_total, created_date FROM bill WHERE bill_id =?`, billid)

	if err != nil {
		return models.Bill{}, "Failed select bill from database", 500
	}

	billModel.Detail = taxList

	return billModel, "", 200
}

func (a Api) GetTaxDetail(taxList []models.Tax) []models.Tax {
	for i,_ := range taxList {
		taxList[i].GetAllDetail()
	}
	return taxList
}

// For dummy data
func (a Api) Dummy() []models.Tax {
	var taxModel []models.Tax
	taxModel = append(taxModel, models.Tax{Name:"Lucky Stretch",TaxCode:2,Price:1000})
	taxModel = append(taxModel, models.Tax{Name:"Big Mac",TaxCode:1,Price:1000})
	taxModel = append(taxModel, models.Tax{Name:"Movie",TaxCode:3,Price:150})
	return taxModel
}