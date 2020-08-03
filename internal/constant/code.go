package constant

// List of tax code.
const (
	FoodCode          int = 1
	TobaccoCode       int = 2
	EntertainmentCode int = 3
)

// TaxCode contains tax code names.
var TaxCode = map[int]string{
	FoodCode:          "Food & Beverage",
	TobaccoCode:       "Tobacco",
	EntertainmentCode: "Entertainment",
}
