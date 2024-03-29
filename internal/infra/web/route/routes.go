package route

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"transaction-system/internal/infra/web/controllers"
)

func Routes(e *echo.Echo, db *sql.DB) {

	accountController := controllers.NewAccountController(db)
	transactionController := controllers.NewTransactionController(db)

	//Accounts
	e.POST("/accounts", accountController.CreateAccount)
	e.GET("/accounts/:accountId", accountController.GetAccountById)

	//Transactions
	e.POST("/transactions", transactionController.CreateTransaction)

	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
