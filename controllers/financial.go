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
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(http.StatusBadRequest, "error", "Request tidak sesuai dengan format json", err.Error()))
		return
	}

	validate := helpers.GetValidator()
	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(http.StatusBadRequest, "error", "Request tidak sesuai dengan aturan validasi", helpers.ParseValidationErrors(err)))
		return
	}

	financial := models.Financial{
		Category: input.Category,
		Nominal: input.Nominal,
		Description: input.Description,
	}

	err = controller.db.Create(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(http.StatusInternalServerError, "error", "Gagal membuat catatan keuangan", err.Error()))
		return
	}
	
	ctx.JSON(http.StatusCreated, helpers.APIResponse(http.StatusCreated, "success", "Berhasil membuat catatan keuangan", financial))
}

func (controller *FinancialController) GetAllFinancial(ctx *gin.Context) {
	category := ctx.Query("category")

	query := controller.db
	if category != "" {
		query = query.Where("category", category)
	}

	var financials []models.Financial
	err := query.Find(&financials).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(http.StatusInternalServerError, "error", "Gagal mengambil data catatan keuangan", err.Error()))
		return
	}
	
	ctx.JSON(http.StatusOK, helpers.APIResponse(http.StatusOK, "success", "Berhasil mengambil data catatan keuangan", financials))
}

func (controller *FinancialController) GetFinancialByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var financial models.Financial
	if err := controller.db.First(&financial, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(http.StatusInternalServerError, "error", "Gagal mengambil data catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponse(http.StatusOK, "success", "Berhasil mengambil data catatan keuangan", financial))
}

func (controller *FinancialController) UpdateFinancial(ctx *gin.Context) {
	id := ctx.Param("id")

	var financial models.Financial
	if err := controller.db.First(&financial, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(http.StatusInternalServerError, "error", "Gagal mengambil data catatan keuangan", err.Error()))
		return
	}

	var input models.FinancialInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(http.StatusBadRequest, "error", "Request tidak sesuai dengan format json", err.Error()))
		return
	}

	validate := helpers.GetValidator()
	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(http.StatusBadRequest, "error", "Request tidak sesuai dengan aturan validasi", helpers.ParseValidationErrors(err)))
		return
	}

	financial.Category = input.Category
	financial.Nominal = input.Nominal
	financial.Description = input.Description

	err = controller.db.Save(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(http.StatusInternalServerError, "error", "Gagal mengubah data catatan keuangan", err.Error()))
		return
	}
	
	ctx.JSON(http.StatusOK, helpers.APIResponse(http.StatusOK, "success", "Berhasil mengubah data catatan keuangan", financial))
}

func (controller *FinancialController) DeleteFinancial(ctx *gin.Context) {
	id := ctx.Param("id")

	var financial models.Financial
	if err := controller.db.Delete(&financial, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(http.StatusInternalServerError, "error", "Gagal menghapus data catatan keuangan", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponse(http.StatusOK, "success", "Berhasil menghapus data catatan keuangan", financial))
}