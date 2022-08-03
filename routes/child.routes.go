package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-mongodb/controllers"
)

type ChildRouteController struct {
	childController controllers.ChildController
}

func NewChildControllerRoute(childController controllers.ChildController) ChildRouteController {
	return ChildRouteController{childController}
}

func (r *ChildRouteController) ChildRoute(rg *gin.RouterGroup) {
	router := rg.Group("/children")

	router.GET("", r.childController.FindChildren)
	router.GET("/:childId", r.childController.FindChildById)
	router.POST("", r.childController.CreateChild)
	router.PATCH("/:childId", r.childController.UpdateChild)
	router.DELETE("/:childId", r.childController.DeleteChild)
}
