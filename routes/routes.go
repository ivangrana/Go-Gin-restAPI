package routes

import (
	"Phinance/controllers"
	"Phinance/handlers"
	"Phinance/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
	}

	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/", controllers.GetAllUsers)
		userGroup.POST("/", controllers.CreateUser)

		user := userGroup.Group("/:id")
		{
			user.GET("/", controllers.GetUserByID)
			user.PUT("/", controllers.UpdateUser)
			user.DELETE("/", controllers.DeleteUser)

			goals := user.Group("/goals")
			goals.Use(middleware.AuthMiddleware())
			{
				goals.GET("/", controllers.GetAllGoals)
				goals.GET("/:goal_id", controllers.GetGoalById)
				goals.POST("/", controllers.CreateGoal)
				goals.PUT("/:goal_id", controllers.UpdateGoal)
				goals.DELETE("/:goal_id", controllers.DeleteGoal)
			}

			budgets := user.Group("/budgets")
			budgets.Use(middleware.AuthMiddleware())
			{
				budgets.GET("/", controllers.GetAllBudgets)
				budgets.GET("/:budget_id", controllers.GetBudgetById)
				budgets.POST("/", controllers.CreateBudget)
				budgets.PUT("/:budget_id", controllers.UpdateBudget)
				budgets.DELETE("/:budget_id", controllers.DeleteBudget)

			}

			transactions := user.Group("/transactions")
			transactions.Use(middleware.AuthMiddleware())
			{
				transactions.GET("/", controllers.GetAllTransactions)
				transactions.GET("/:transaction_id", controllers.GetTransactionById)
				transactions.POST("/", controllers.CreateTransaction)
				transactions.DELETE("/:transaction_id", controllers.DeleteTransaction)
			}

			receipts := user.Group("/receipts")
			receipts.Use(middleware.AuthMiddleware())
			{
				receipts.GET("/", controllers.GetAllReceipts)
				receipts.GET("/:receipt_id", controllers.GetReceiptById)
				receipts.POST("/", controllers.CreateReceipt)
				receipts.PUT("/:receipt_id", controllers.UpdateReceipt)
				receipts.DELETE("/:receipt_id", controllers.DeleteReceipt)
			}
		}
	}

	categories := router.Group("/categories")

	{
		categories.GET("/", controllers.GetAllCategories)
		categories.POST("/", controllers.CreateCategory)
		categories.GET("/:category_id", controllers.GetCategoryById)
		categories.PUT("/:category_id", controllers.UpdateCategory)
		categories.DELETE("/:category_id", controllers.DeleteCategory)
	}

	marketProductGroup := router.Group("/market-products")
	{
		marketProductGroup.GET("/", controllers.GetAllProducts)
		marketProductGroup.POST("/", controllers.CreateMarketProduct)
		marketProductGroup.GET("/:product_id", controllers.GetMarketProductById)
		marketProductGroup.PUT("/:product_id", controllers.UpdateMarketProduct)
		marketProductGroup.DELETE("/:product_id", controllers.DeleteMarketProduct)
	}

}
