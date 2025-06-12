package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Mafilala/ketero/backend/controllers"
)
func RegisterMeasure(r *gin.Engine) {
   measureGroup := r.Group("/measure")
    {
        measureGroup.POST("/", controllers.CreateMeasure)
        measureGroup.DELETE("/:id", controllers.DeleteMeasure)
        measureGroup.GET("/:id", controllers.GetMeasureByID)
        measureGroup.GET("/", controllers.GetAllMeasures)
	measureGroup.PUT("/:id", controllers.UpdateMeasure)


    }
}

func RegisterClothingType(r *gin.Engine) {
    clothingTypeGroup := r.Group("/clothingType")
    {
        clothingTypeGroup.POST("/", controllers.CreateClothingType)
        clothingTypeGroup.DELETE("/:id", controllers.DeleteClothingType)
        clothingTypeGroup.GET("/", controllers.GetAllClothingTypes)
        clothingTypeGroup.GET("/:id", controllers.GetClothingTypeByID)
	clothingTypeGroup.PUT("/:id", controllers.UpdateClothingType)

    }
}

func RegisterClothing(r *gin.Engine) {
	clothingGroup := r.Group("/clothing")
	{
		clothingGroup.POST("/", controllers.CreateClothing)
		clothingGroup.PUT("/:id", controllers.UpdateClothing)
		clothingGroup.DELETE("/:id", controllers.DeleteClothing)
		clothingGroup.GET("/", controllers.GetAllClothing)
		clothingGroup.GET("/:id", controllers.GetClothingByID)
	}
}

func RegisterClient(r *gin.Engine) {
	clientGroup := r.Group("/client")
	{
		clientGroup.POST("/", controllers.CreateClient)
		clientGroup.DELETE("/:id", controllers.DeleteClient)
		clientGroup.GET("/", controllers.GetAllClients)
		clientGroup.GET("/:id", controllers.GetClientByID)
	}
}

func RegisterStatus(r *gin.Engine) {
    statusGroup := r.Group("/status")
    {
        statusGroup.POST("/", controllers.CreateStatus)
        statusGroup.DELETE("/:id", controllers.DeleteStatus)
        statusGroup.GET("/", controllers.GetAllStatuses)
        statusGroup.GET("/:id", controllers.GetStatusByID)
	statusGroup.PUT("/:id", controllers.UpdateStatus)
    }
}

func RegisterAddClothing(r *gin.Engine) {
    clothingPartGroup := r.Group("/clothing_part")
    {
        clothingPartGroup.POST("/", controllers.AddClothing)
        clothingPartGroup.DELETE("/:clothing_type_id/:clothing_id", controllers.RemoveClothingTypePart)
        clothingPartGroup.GET("/:id", controllers.GetAllClothingParts)    
    }
}

func RegisterClothingMeasures(r *gin.Engine) {
    clothingMeasurePartGroup := r.Group("/clothing_measure")
    {
        clothingMeasurePartGroup.POST("/", controllers.AddMeasure)
        clothingMeasurePartGroup.DELETE("/:clothing_id/:measure_id", controllers.RemoveMeasure)
        clothingMeasurePartGroup.GET("/:id", controllers.GetAllClothingMeasures)    
    }
}

func RegisterOrderRoutes(r *gin.Engine) {
    orderGroup := r.Group("/order")
    {
        orderGroup.POST("/", controllers.CreateOrder)
        orderGroup.GET("/:id", controllers.GetOrderByID)
        orderGroup.DELETE("/:id", controllers.DeleteOrder)
        orderGroup.GET("/", controllers.GetAllOrders)
	orderGroup.PATCH("/:id", controllers.PatchOrder)

    }
}

func RegisterOrderMeasureRoutes(r *gin.Engine) {
	orderMeasureGroup := r.Group("/order-measure")
	{
		orderMeasureGroup.POST("/", controllers.CreateOrderMeasure)
		orderMeasureGroup.PUT("/:order_id", controllers.UpdateOrderMeasure)
		orderMeasureGroup.DELETE("/:order_id/:measure_id", controllers.DeleteOrderMeasure)
		orderMeasureGroup.GET("/:order_id", controllers.GetOrderMeasuresByOrderID)
	}
}

func RegisterPriceDetailRoutes(r *gin.Engine) {
	group := r.Group("/price-detail")
	{
		group.POST("/", controllers.CreatePriceDetail)
		group.GET("/:order_id", controllers.GetPriceDetailByOrderID)
		group.PUT("/:order_id", controllers.UpdatePriceDetail)
		group.DELETE("/:order_id", controllers.DeletePriceDetail)
	}
}

func RegisterOrderDetailRoutes(r *gin.Engine) {
	orderDetailGroup := r.Group("/order-detail")
	{
		orderDetailGroup.POST("/", controllers.CreateOrderDetail)
		orderDetailGroup.GET("/:order_id", controllers.GetOrderDetail)
		orderDetailGroup.PUT("/:order_id", controllers.UpdateOrderDetail)
		orderDetailGroup.DELETE("/:order_id", controllers.DeleteOrderDetail)
	}
}

