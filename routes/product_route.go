package routes

import (
	"ghn-test/controllers"
	"ghn-test/middlewares"
	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Echo) {
	e.GET("/jwt", controllers.GetJWT)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:productId", controllers.EditProduct)
	e.GET("/products/:productId", controllers.GetProduct)
	e.GET("/products", controllers.GetAllProducts, middlewares.Authorization())
	e.DELETE("/products/:productId", controllers.DeleteProduct)
}
