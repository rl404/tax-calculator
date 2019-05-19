package controllers

import (
	"github.com/revel/revel"
	"encoding/json"
	"strconv"
	"fmt"

	"github.com/rl404/tax-calculator/app/models"
	"github.com/rl404/tax-calculator/app/helpers"
)

type Tax struct {
	*revel.Controller
}

var createBillUrl = "http://localhost:9001/api/createbill"
var getBillUrl = "http://localhost:9001/api/GetBill"

var moreStyles = []string {"css/tax.css"}
var moreScripts = []string {"js/tax.js", "js/simple-money-format.js"}

func (c Tax) Index() revel.Result {
	return c.Render(moreStyles)
}

func (c Tax) Create() revel.Result {

	// Get tax type list for dropdown
	taxTypeList := models.TaxDetail{}.GetDetailList()

	postData := c.Params.Form
	if len(postData) != 0 {
		var taxList []models.Tax
		for i, _ := range postData["name"] {
			if postData["name"][i] == "" {
				continue
			}
			postTaxCode, _ := strconv.Atoi(postData["taxcode"][i])
			postPrice, _ := strconv.Atoi(postData["price"][i])
			var taxModel = models.Tax{
				Name 	: postData["name"][i],
				TaxCode : postTaxCode,
				Price 	: float64(postPrice),
			}
			taxList = append(taxList, taxModel)
		}

		// Convert to json and send
		taxListJson, _ := json.Marshal(taxList)
		_, responseData := helpers.JsonPost(createBillUrl, taxListJson)

		var responseArr map[string]interface{}
		err := json.Unmarshal([]byte(responseData), &responseArr)
		if err != nil {
			panic(err)
		}

		if int(responseArr["status"].(float64)) != 200 {
			hasError := responseArr["message"]
			return c.Render(hasError, taxTypeList, moreStyles, moreScripts)
		}

		newBill := responseArr["data"].(map[string]interface{})["billid"]
		return c.Redirect("/tax/result?bill=%d", int(newBill.(float64)))
	}

	return c.Render(taxTypeList, moreStyles, moreScripts)
}

func (c Tax) Result() revel.Result {

	billId := c.Params.Query.Get("bill")

	_, responseData := helpers.SendGetRequest(fmt.Sprintf("%s?bill=%s", getBillUrl, billId))

	var responseArr map[string]interface{}
	err := json.Unmarshal([]byte(responseData), &responseArr)
	if err != nil {
		panic(err)
	}

	if int(responseArr["status"].(float64)) != 200 {
		hasError := responseArr["message"]
		return c.Render(hasError, moreStyles)
	}

	billModel := responseArr["data"]
	return c.Render(billModel, moreStyles)
}

func (c Tax) Find() revel.Result {

	postData := c.Params.Form
	if len(postData) != 0 {
		return c.Redirect("/tax/result?bill=%s", postData["bill"][0])
	}

	return c.Render(moreStyles)
}