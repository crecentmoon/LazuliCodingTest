package lazuli

import (
	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
	"github.com/labstack/echo"
)

func InitServer(db repository.SqlHandler) {

	// infra.NewSqlHandler().InitRDB()

	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
