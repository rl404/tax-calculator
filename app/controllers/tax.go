package controllers

import (
	"github.com/revel/revel"
	"github.com/rl404/tax-calculator/app/models"
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