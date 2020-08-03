package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/rl404/tax-calculator/internal/config"
	"github.com/rl404/tax-calculator/internal/constant"
	"github.com/rl404/tax-calculator/internal/model"
	"github.com/rl404/tax-calculator/internal/view"
)

// TaxHandler to handle all tax activities.
type TaxHandler struct {
	DB *gorm.DB
}

// newTaxHandler to create new instance.
func newTaxHandler(cfg config.Config) (th TaxHandler, err error) {
	// Init db connection.
	th.DB, err = cfg.InitDB()
	if err != nil {
		return th, err
	}

	return th, nil
}

// getList to get user's tax list.
func (th *TaxHandler) getList(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	if userID == 0 {
		view.RespondWithJSON(w, http.StatusBadRequest, constant.ErrRequiredUser.Error(), nil)
		return
	}

	var userItems []model.Item
	th.DB.Where("user_id = ?", userID).Find(&userItems)

	for i, j := range userItems {
		th.fillTax(&j)
		userItems[i] = j
	}

	view.RespondWithJSON(w, http.StatusOK, http.StatusText(http.StatusOK), userItems)
}

// add to add a new tax.
func (th *TaxHandler) add(w http.ResponseWriter, r *http.Request) {
	var request model.Item

	// Get request body.
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		view.RespondWithJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Validate request data.
	err = th.validateRequest(request)
	if err != nil {
		view.RespondWithJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Insert to db.
	err = th.DB.Create(&request).Error
	if err != nil {
		view.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Calculate and fill tax field.
	th.fillTax(&request)

	view.RespondWithJSON(w, http.StatusCreated, http.StatusText(http.StatusCreated), request)
}

// delete to delete tax item.
func (th *TaxHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 {
		view.RespondWithJSON(w, http.StatusBadRequest, constant.ErrRequiredID.Error(), nil)
		return
	}

	var userItem model.Item
	if th.DB.Where("id = ?", id).First(&userItem).RecordNotFound() {
		view.RespondWithJSON(w, http.StatusNotFound, constant.ErrNotFound.Error(), nil)
		return
	}

	err := th.DB.Delete(&userItem).Error
	if err != nil {
		view.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	view.RespondWithJSON(w, http.StatusOK, http.StatusText(http.StatusOK), nil)
}

// validateRequest to validate request data.
func (th *TaxHandler) validateRequest(request model.Item) error {
	if request.UserID == 0 {
		return constant.ErrRequiredUser
	}

	if request.Name == "" {
		return constant.ErrRequiredName
	}

	if constant.TaxCode[request.TaxCode] == "" {
		return constant.ErrInvalidTaxCode
	}

	if request.Price <= 0 {
		return constant.ErrInvalidPrice
	}

	return nil
}

// fillTax to calculate and fill tax field.
func (th *TaxHandler) fillTax(item *model.Item) {
	item.Type = constant.TaxCode[item.TaxCode]

	if item.TaxCode == constant.FoodCode {
		item.Refundable = true
		item.Tax = item.Price / 10.0
	} else if item.TaxCode == constant.TobaccoCode {
		item.Refundable = false
		item.Tax = 10 + item.Price/50.0
	} else if item.TaxCode == constant.EntertainmentCode {
		item.Refundable = false
		if item.Price >= 100 {
			item.Tax = (item.Price - 100.0) / 100.0
		}
	}

	item.Amount = item.Price + item.Tax
}
