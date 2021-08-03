
package controllers

import (
    "strconv"
    "fmt"
    "time"
    "net/http"
    "github.com/labstack/echo"
    "app/models"
    "app/database"
)

func InternalError(c echo.Context) error {
    res_map := make(map[string]string, 2)
    res_map["message"] = "internal server error"
    return c.JSON(http.StatusInternalServerError, &res_map)
}

type RestaurantController struct{}

func NewRestaurantController() *RestaurantController {
    return new(RestaurantController)
}

func (hc *RestaurantController) Index(c echo.Context) error {
    return c.JSON(http.StatusOK, "a")
}

func (hc *RestaurantController) GetRestaurantList(c echo.Context) error {
    var restaurants []models.Restaurant
    db := database.GetDB()
    if err := db.Find(&restaurants).Error; err != nil {
        fmt.Println("fail to get restaurant list",)
        return c.JSON(http.StatusNotFound, "a")
    }

    return c.JSON(http.StatusOK, &restaurants)
}

func (hc *RestaurantController) GetRestaurantByID(c echo.Context) error {
    id := c.Param("id")

    id_u64, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        fmt.Println("fail to getting restaurant id", id)
        return c.JSON(http.StatusBadRequest, "fail to parse id")
    }
    
    var restaurant models.Restaurant
    if err := restaurant.FindByID(id_u64); err != nil {
        fmt.Println("fail to getting restaurant id",)
        return c.JSON(http.StatusNotFound, "a")
    }

    // json_res, err := json.Marshal(&restaurant)
    if err != nil {
        fmt.Println(err)
        return InternalError(c)
    }
    return c.JSON(http.StatusOK, &restaurant)
}


func (hc *RestaurantController) Create(c echo.Context) error {
    // json_map := make(map[string]interface{})
    // err := json.NewDecoder(c.Request().Body).Decode(&json_map)
    // if err != nil {
    //     return err
    // } else {
    //     //json_map has the JSON Payload decoded into a map
    //     cb_type := json_map["type"]
    //     challenge := json_map["challenge"]
    //     fmt.Println(cb_type)
    //     fmt.Println(challenge)
    // }


    fmt.Println("request body:", c.Request().Body)
    fmt.Println(c.Request().Header.Get("Content-Type"))
    // restaurant := new(models.Restaurant)
    var restaurant models.Restaurant;
    if err := c.Bind(&restaurant); err != nil {
        fmt.Println("fail bind request to restaurant")
        fmt.Println(err)
        return InternalError(c)
    }
    restaurant.CreatedAt = time.Now()
    restaurant.UpdatedAt = time.Now()
    fmt.Println("restaurant: ", restaurant)
    db := database.GetDB()
    result := db.Create(&restaurant)
    if err := result.Error; err != nil {
        fmt.Println("fail create restaurant: ", restaurant.ID, restaurant.Name)
        fmt.Println(err)
        return InternalError(c)
    }
    return c.JSON(http.StatusOK, "create " + restaurant.Name)
}