package handler

// func Init() {
// 	e := echo.New()

// 	e.Use(middleware.Recover())
// 	e.Use(middleware.Logger())

// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins:     []string{"*"},
// 		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
// 		AllowCredentials: true,
// 		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
// 	}))

// 	firebase, err := NewFireBaseHandler()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	storage, err := NewStorageHandler()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	userController := controllers.NewUserController(NewSqlHandler(), firebase, storage, NewCustomValidator())

// 	api := e.Group("api")

// 	// ユーザー
// 	api.GET("/getUser", userController.GetUser)
// 	api.POST("/createUser", userController.CreateUser)
// 	api.POST("/login", userController.Login)
// 	api.POST("/user/login/google", userController.GoogleLogin)
// 	api.PUT("/user/update/basicInfo", userController.UpdateBasicData)
// 	api.PATCH("/user/name", userController.UpdateUserName)
// 	api.PATCH("/user/email", userController.UpdateEmail)
// 	api.PATCH("/user/phoneNumber", userController.UpdatePhoneNumber)
// 	api.DELETE("/user/delete", userController.Delete)

// 	e.Logger.Fatal(e.Start(":8080"))
// }
