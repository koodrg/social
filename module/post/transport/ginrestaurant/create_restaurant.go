package ginrestaurant

import (
	"net/http"
	"social/common"
	restaurantmodel "social/module/post/model"
	restaurantstorage "social/module/post/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(db)
		if err := store.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
