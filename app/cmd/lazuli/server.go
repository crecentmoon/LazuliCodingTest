package lazuli

import (
	"github.com/labstack/echo"
)

func InitServer() {

	// infra.NewSqlHandler().InitRDB()

	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
