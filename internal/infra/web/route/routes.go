package route

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"transaction-system/internal/infra/web/controllers"
)

func Routes(e *echo.Echo, db *sql.DB) {

	accountController := controllers.NewAccountController(db)

	//Accounts
	e.POST("/accounts", accountController.CreateAccount)

	//Transactions
}
