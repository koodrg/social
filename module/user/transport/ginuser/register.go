package ginuser

import (
	"net/http"
	"social/common"
	usermodel "social/module/user/model"
	userstorage "social/module/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSqlStore(db)
		if err := store.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
