package main

import (
	"fmt"
	"os"

	restaurantgin "example.com/g07-food-delivery/modules/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("db %T, err %T", db, err)
	db = db.Debug()

	/// -------------- Insert DB --------------
	// newRestaurant := restaurantmodel.Restaurant{Name: "200Lab", Address: "Some where"}
	// if err := db.Create(&newRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	/// -------------- Find one --------------
	// var oldRes restaurantmodel.Restaurant
	// // if err := db.Where("id = ?", 1).First(&oldRes).Error; err != nil {
	// // 	log.Println(err)
	// // }
	// if err := db.Where(map[string]interface{}{"id": 1}).First(&oldRes).Error; err != nil { /// co the nhan vao 1 map interface
	// 	log.Println(err)
	// }
	// log.Println(oldRes)

	/// -------------- Find list --------------
	// var listRes []restaurantmodel.Restaurant
	// if err := db.Limit(10).Find(&listRes).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(listRes)

	/// -------------- Update --------------
	// emptyName := ""
	// dataUpdate := restaurantmodel.RestaurantCreate{
	// 	Name: &emptyName, /// Chuoi rong la 1 bien co value => update duoc
	// 	/// Address dang la nil
	// }
	// if err := db.Where("id = ?", 1).Updates(&dataUpdate).Error; err != nil {
	// 	log.Println(err)
	// }

	/// -------------- Delte --------------
	// if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		// Create restaurant
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurant(db))

			// Get restaurant detail
			restaurants.GET("/:id", restaurantgin.GetRestaurant(db))

			// Get restaurant list
			restaurants.GET("", restaurantgin.ListRestaurant(db))

			// Update restaurant
			restaurants.PUT("/:id", restaurantgin.UpdateRestaurant(db))

			// Update restaurant
			restaurants.DELETE("/:id", restaurantgin.DeleteRestaurant(db))
		}
	}

	r.Run()
}
