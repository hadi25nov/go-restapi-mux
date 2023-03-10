package productcontroller

import (
	"net/http"

	"github.com/hadi25nov/go-restapi-mux/helper"

	"github.com/hadi25nov/go-restapi-mux/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {

	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, products)

}
func Show(w http.ResponseWriter, r *http.Request) {
	// var product
	// vars := mux.Vars(r)
	// id.err := strconv

}
func Create(w http.ResponseWriter, r *http.Request) {

}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
