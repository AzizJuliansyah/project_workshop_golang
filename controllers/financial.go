package controllers

import (
	"net/http"
	"project_workshop_golang_test/helpers"
	"project_workshop_golang_test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FinancialController struct {
	db *gorm.DB
}

func NewFinancialController(db *gorm.DB) FinancialController {
	return FinancialController{db: db}
}

func (controller *FinancialController) CreateFinancial(ctx *gin.Context) {
	var input models.FinancialInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponseFormat(http.StatusBadRequest, "error", "Format JSON tidak sesuai", err.Error()))
		return
	}

	validate := helpers.GetValidator()
	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponseFormat(http.StatusBadRequest, "error", "Error validation", helpers.ParseValidationErrors(err)))
		return
	} 

	financial := models.Financial{
		Category: input.Category,
		Nominal: input.Nominal,
		Information: input.Information,
	}

	err = controller.db.Create(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponseFormat(http.StatusInternalServerError, "error", "Gagal membuat catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.APIResponseFormat(http.StatusCreated, "success", "Berhasil menambahkan catatan keuangan", financial))
}

func (controller *FinancialController) GetAllFinancial(ctx *gin.Context) {
	category := ctx.Query("category")

	var financial []models.Financial
	query := controller.db
	if category != "" {
		query = query.Where("category = ?", category)
	}


	err := query.Find(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponseFormat(http.StatusInternalServerError, "error", "Gagal mengambil catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponseFormat(http.StatusOK, "success", "List catatan keuangan", financial))
}

func (controller *FinancialController) GetFinancialByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var financial models.Financial
	err := controller.db.First(&financial, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponseFormat(http.StatusInternalServerError, "error", "Gagal mengambil catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponseFormat(http.StatusOK, "success", "Catatan keuangan", financial))
}

func (controller *FinancialController) UpdateFinancial(ctx *gin.Context) {
	id := ctx.Param("id")

	var input models.FinancialInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponseFormat(http.StatusBadRequest, "error", "Format JSON tidak sesuai", err.Error()))
		return
	}
	
	validate := helpers.GetValidator()
	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponseFormat(http.StatusBadRequest, "error", "Error validation", helpers.ParseValidationErrors(err)))
		return
	} 

	var financial models.Financial
	err = controller.db.First(&financial, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponseFormat(http.StatusInternalServerError, "error", "Gagal mengambil catatan keuangan", err.Error()))
		return
	}

	financial.Category = input.Category
	financial.Nominal = input.Nominal
	financial.Information = input.Information

	err = controller.db.Save(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponseFormat(http.StatusInternalServerError, "error", "Gagal membuat catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.APIResponseFormat(http.StatusCreated, "success", "Berhasil mengubah data catatan keuangan", financial))
}

func (controller *FinancialController) DeleteFinancial(ctx *gin.Context) {
	id := ctx.Param("id")

	var financial models.Financial
	err := controller.db.Delete(&financial, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponseFormat(http.StatusInternalServerError, "error", "Gagal mengambil catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponseFormat(http.StatusOK, "success", "Berhasil menghapus data catatan keuangan", nil))
}