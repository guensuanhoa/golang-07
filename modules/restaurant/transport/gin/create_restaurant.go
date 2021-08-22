package restaurantgin

import (
	"net/http"

	restaurantbiz "example.com/g07-food-delivery/modules/restaurant/biz"
	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
	restaurantstorage "example.com/g07-food-delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantmodel.RestaurantCreate

		// Nen dung shouldBind vi no co return error de check co bind duoc khong
		if err := c.ShouldBind(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Dependency injection
		// Trong bat ky kien truc nao cung phai co mot noi nhu the nay
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantStore(store)

		// Save to DB
		if err := biz.CreateNewRestaurant(c.Request.Context(), &newData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// Return restaurant id
		c.JSON(http.StatusOK, gin.H{"data": newData.Id})
	}
}
