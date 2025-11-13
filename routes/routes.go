package routes

import (
	"project_workshop_golang_test/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	apiVersion := router.Group("/api/v1")


	financialControllers := controllers.NewFinancialController(db)
	financialGroup := apiVersion.Group("/financial")
	{
		financialGroup.POST("", financialControllers.CreateFinancial)
		financialGroup.GET("", financialControllers.GetAllFinancial)
		financialGroup.GET("/:id", financialControllers.GetFinancialByID)
		financialGroup.PUT("/:id", financialControllers.UpdateFinancial)
		financialGroup.DELETE("/:id", financialControllers.DeleteFinancial)
	}
	
}