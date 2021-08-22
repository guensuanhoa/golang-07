package restaurantgin

import (
	"net/http"

	"example.com/g07-food-delivery/common"
	restaurantbiz "example.com/g07-food-delivery/modules/restaurant/biz"
	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
	restaurantstorage "example.com/g07-food-delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		// Verify query
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := paging.Process(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.ListRestaurantStore(store)
		// Find id in DB
		result, err := biz.ListDataWithCondition(c.Request.Context(), &paging, &filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return restaurant json
		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
