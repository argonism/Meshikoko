package server

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "app/config"
    "app/controllers"
    "fmt"
)


func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
    fmt.Printf("Request Body: %v\n", string(reqBody))
    fmt.Printf("Response Body: %v\n", string(resBody))
  }

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
    c := config.GetConfig()
    router := echo.New()
    router.Use(middleware.BodyDump(bodyDumpHandler))
    router.Use(middleware.Logger())
    router.Use(middleware.Recover())
    router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: c.GetStringSlice("server.cors"),
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
    }))

    version := router.Group("/" + c.GetString("server.version"))

    restaurantController := controllers.NewRestaurantController()
    version.GET("/restaurant", restaurantController.Index)
    version.GET("/restaurants", restaurantController.GetRestaurantList)
    version.GET("/restaurant/:id", restaurantController.GetRestaurantByID)
    version.POST("/restaurant", restaurantController.Create)

    return router, nil
}