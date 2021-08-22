package restaurantgin

import (
	"net/http"
	"strconv"

	restaurantbiz "example.com/g07-food-delivery/modules/restaurant/biz"
	restaurantstorage "example.com/g07-food-delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Verify restaurant id
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Dependency injection
		// Trong bat ky kien truc nao cung phai co mot noi nhu the nay
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		// Find id in DB
		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return restaurant json
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
