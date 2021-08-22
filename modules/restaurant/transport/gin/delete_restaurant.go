package restaurantgin

import (
	"fmt"
	"net/http"
	"strconv"

	restaurantbiz "example.com/g07-food-delivery/modules/restaurant/biz"
	restaurantstorage "example.com/g07-food-delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Verify restaurant id
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Deleting id: %v", id)

		// Dependency injection
		// Trong bat ky kien truc nao cung phai co mot noi nhu the nay
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		// Update restaurant in DB
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return status success
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
